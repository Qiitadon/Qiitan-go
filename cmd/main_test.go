package main

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/Qithub-BOT/Qiitan-go/qiitan/compiler"
	"github.com/d5/tengo/v2/require"
	"github.com/stretchr/testify/assert"
	"github.com/zenizh/go-capturer"
)

// 正常系・基本的な Qiitan スクリプトの実行（コンパイルと実行の）テスト。
func TestRun(t *testing.T) {
	// コマンド引数（os.Args）のバックアップと遅延リカバリ
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// テスト用スクリプトのパス取得
	pathFileScript, err := filepath.Abs("../examples/hello_world/hello_world.qiitan")
	require.NoError(t, err)

	out := capturer.CaptureStdout(func() {
		// ダミーのコマンド引数をセット
		os.Args = []string{
			t.Name(),
			pathFileScript,
		}

		main() // 実行
	})

	// テスト結果のアサーション
	assert.Contains(t, out, "Hello World!\n")
}

// コンパイルのみ（バイナリ変換）のテスト。
func TestRun_compile(t *testing.T) {
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// テスト用スクリプトのパス取得
	pathFileScript, err := filepath.Abs("../examples/hello_world/hello_world.qiitan")
	require.NoError(t, err)

	// コンパイル済みファイルの出力先
	pathFileBin := filepath.Join(t.TempDir(), util.RandStr(8))

	// 実行と出力のキャプチャ
	out := capturer.CaptureStdout(func() {
		os.Args = []string{
			t.Name(),
			"-o", pathFileBin,
			pathFileScript,
		}

		main()
	})

	assert.Contains(t, out, "TestRun_compile")
	assert.NotContains(t, out, "Hello World", "-o option should not execute the script")
	assert.FileExists(t, pathFileBin)
}

// 引数なしテスト。
func TestRun_no_args(t *testing.T) {
	// コマンド引数（os.Args）のバックアップと遅延リカバリ
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	out := capturer.CaptureStdout(func() {
		// ダミーのコマンド引数をセット
		os.Args = []string{
			t.Name(),
		}

		main() // 実行
	})

	// テスト結果のアサーション
	assert.Contains(t, out, "Welcome to Qiitan Script Interactive Mode")
	assert.Contains(t, out, "> ")
}

// 構文エラー付きスクリプトのテスト。
func TestRun_no_scriptfiles(t *testing.T) {
	// コマンド引数（os.Args）のバックアップと遅延リカバリ
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// テスト用スクリプトのパス取得
	pathFileScript, err := filepath.Abs("../testdata/malformed.qiitan")
	require.NoError(t, err)

	out := capturer.CaptureStdout(func() {
		// ダミーのコマンド引数をセット
		os.Args = []string{
			t.Name(),
			pathFileScript,
		}

		err := Run() // 実行

		require.Error(t, err)
	})

	// 標準出力のアサーション
	assert.Empty(t, out)
}

// コンパイル時の --no-resolve オプション付きのテスト。
func TestRun_do_not_resolve_path_option(t *testing.T) {
	// コマンド引数（os.Args）のバックアップと遅延リカバリ
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// テスト用スクリプトのパス取得
	pathFileScript, err := filepath.Abs("../examples/hello_world/hello_world.qiitan")
	require.NoError(t, err)

	out := capturer.CaptureStdout(func() {
		// ダミーのコマンド引数をセット
		os.Args = []string{
			t.Name(),
			"--no-resolve", // 相対パスの名前解決（tengo v3 で削除予定らしい）
			pathFileScript,
		}

		main() // 実行
	})

	// テスト結果のアサーション
	assert.Contains(t, out, "Hello World!\n")
}

// バイトコード出力先のファイルオープンにエラーが発生した場合のエラーメッセージ
// 確認テスト。
func TestRun_file_exists(t *testing.T) {
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	oldOsOpenFile := compiler.OsOpenFile
	defer func() {
		compiler.OsOpenFile = oldOsOpenFile
	}()

	compiler.OsOpenFile = func(name string, flag int, perm fs.FileMode) (*os.File, error) {
		return nil, errors.New("dummy: failed to open file")
	}

	// テスト用スクリプトのパス取得
	pathFileScript, err := filepath.Abs("../examples/hello_world/hello_world.qiitan")
	require.NoError(t, err)

	a, err := os.CreateTemp(t.TempDir(), "")
	require.NoError(t, err)

	defer a.Close()

	pathFileOut := a.Name() + ".qtn"
	t.Logf("out file path: %v", pathFileOut)

	out := capturer.CaptureOutput(func() {
		os.Args = []string{
			t.Name(),
			"-o", pathFileOut,
			pathFileScript,
		}

		err := Run()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "dummy: failed to open file")
	})

	assert.Empty(t, out, "on error it should not print anything")
}

// コンパイル済みのファイルを実行するテスト。
func TestRun_run_compiled(t *testing.T) {
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// テスト用コンパイル済みファイル
	pathFileBin, err := filepath.Abs("../testdata/hello_world")
	require.NoError(t, err)

	out := capturer.CaptureStdout(func() {
		os.Args = []string{
			t.Name(),
			pathFileBin,
		}

		main()
	})

	assert.Equal(t, "Hello World!\n", out)
}

// シバン付きの Qiitan スクリプトを実行するテスト。
func TestRun_shebang(t *testing.T) {
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// テスト用コンパイル済みファイル
	pathFileBin, err := filepath.Abs("../examples/shebang/shebang.qiitan")
	require.NoError(t, err)

	out := capturer.CaptureStdout(func() {
		os.Args = []string{
			t.Name(),
			pathFileBin,
		}

		main()
	})

	assert.Equal(t, "Hello World!\n", out)
}

// ヘルプ表示のテスト。
func TestRun_show_help(t *testing.T) {
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	out := capturer.CaptureStdout(func() {
		os.Args = []string{
			t.Name(),
			"--help",
		}

		main()
	})

	assert.Contains(t, out, "Usage:\n")
}

// バージョン情報表示のテスト。
func TestRun_show_version(t *testing.T) {
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	out := capturer.CaptureStdout(func() {
		os.Args = []string{
			t.Name(),
			"--version",
		}

		main()
	})

	assert.Contains(t, out, "dev\n")
}

// 未定義フラグ／オプションのテスト。
func TestRun_wrong_flag(t *testing.T) {
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	out := capturer.CaptureStderr(func() {
		os.Args = []string{
			t.Name(),
			"--foobarbuzz",
		}

		err := Run()

		require.Error(t, err, "undefined flag should return an error")
	})

	assert.Contains(t, out, "flag provided but not defined")
}
