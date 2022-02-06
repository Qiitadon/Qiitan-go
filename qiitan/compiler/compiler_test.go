package compiler_test

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/Qiitadon/Qiitan-go/qiitan/compiler"
	"github.com/Qiitadon/Qiitan-go/qiitan/modules"
	"github.com/d5/tengo/v2/require"
	"github.com/stretchr/testify/assert"
	"github.com/zenizh/go-capturer"
)

// ============================================================================
//  Examples
// ============================================================================

func ExampleNew() {
	// コンパイラの生成
	cpl := compiler.New()

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	// スクリプト
	script := []byte(`
		foo := import("fmt")
		foo.println("Hello World!")
	`)

	// コンパイルと実行
	err := cpl.CompileAndRun(mods, script, "")
	if err != nil {
		log.Fatal(err)
	}

	// Output: Hello World!
}

// ============================================================================
//  Method Tests
// ============================================================================

// ----------------------------------------------------------------------------
//  CompileAndRun()
// ----------------------------------------------------------------------------

// 構造自体は問題ない（パースはできる）が、構文に問題ありで実行（コンパイル）を
// 失敗させるテスト。
func TestCompileAndRun_error_compile(t *testing.T) {
	// コンパイラの生成
	cpl := compiler.New()

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	// 不正なスクリプト
	script := []byte(`func() { break }`)

	out := capturer.CaptureOutput(func() {
		// コンパイルと実行
		err := cpl.CompileAndRun(mods, script, "")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "break not allowed outside loop")
	})

	assert.Empty(t, out)
}

// ----------------------------------------------------------------------------
//  CompileOnly()
// ----------------------------------------------------------------------------

// 正常系ファイルのコンパイル済みファイル出力テスト。
func TestCompileOnly_out_file(t *testing.T) {
	// テスト用スクリプトの読み込みパス取得
	nameFileScript := "hello_world.qiitan"
	pathFileOrigin := getPathExample(t, "hello_world", nameFileScript)

	// バイトデータの出力ファイル名（期待するファイル名）
	nameFileOut := "hello_world.qtn"

	// コンパイラの生成
	cpl := compiler.New()

	// qiitan スクリプトのデータ読み込み
	byteData, err := cpl.ReadScript(pathFileOrigin)
	require.NoError(t, err)

	// カレント・ディレクトリを temp に移動
	pathDirTmp := t.TempDir()

	oldPWD := util.ChDir(pathDirTmp)
	defer oldPWD() // 処理後カレント・ディレクトリを元に戻す

	// テスト用スクリプトを temp にコピー
	pathFileScript := filepath.Join(pathDirTmp, nameFileScript)
	require.NoError(t, os.WriteFile(pathFileScript, byteData, os.ModePerm))

	// コンパイル済みファイルの出力先パス取得
	expectPathFileOut := filepath.Join(pathDirTmp, nameFileOut)

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	// 実行と標準出力・標準エラー出力のキャプチャ
	out := capturer.CaptureOutput(func() {
		// temp ディレクトリ先のスクリプト実行
		err := cpl.CompileOnly(mods, byteData, pathFileScript, "")

		require.NoError(t, err)
	})

	assert.FileExists(t, expectPathFileOut, "compiled file should output to current dir")
	assert.FileExists(t, strings.TrimSpace(out), "the output should contain the file path")
}

// Markdown（Qiitan スクリプトでない）ファイルのコンパイルエラー・テスト。
func TestCompileOnly_bad_script(t *testing.T) {
	cpl := compiler.New()

	// テスト用スクリプトのパス取得
	pathFileScript := getPathTestData(t, "malformed.qiitan")

	// コンパイルされてないデータの読み込み
	byteData, err := cpl.ReadScript(pathFileScript)
	require.NoError(t, err)

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	// 実行と標準出力・標準エラー出力のキャプチャ
	out := capturer.CaptureOutput(func() {
		err := cpl.CompileOnly(mods, byteData, pathFileScript, "")

		require.Error(t, err)
		assert.Contains(t, err.Error(),
			"failed to parse file: Parse Error: illegal character U+0023")
	})

	assert.Empty(t, out)
}

// バイトコード出力先のファイルオープンに失敗した時のテスト。
func TestCompileOnly_fail_to_open_outfile(t *testing.T) {
	// テスト用スクリプトの読み込みパス取得
	nameFileScript := "hello_world.qiitan"
	pathFileOrigin := getPathExample(t, "hello_world", nameFileScript)

	// コンパイラの生成
	cpl := compiler.New()

	// qiitan スクリプトのデータ読み込み
	byteData, err := cpl.ReadScript(pathFileOrigin)
	require.NoError(t, err)

	// カレント・ディレクトリを temp に移動
	pathDirTmp := t.TempDir()

	oldPWD := util.ChDir(pathDirTmp)
	defer oldPWD() // 処理後カレント・ディレクトリを元に戻す

	// テスト用スクリプトを temp にコピー
	pathFileScript := filepath.Join(pathDirTmp, nameFileScript)
	require.NoError(t, os.WriteFile(pathFileScript, byteData, os.ModePerm))

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	// ファイル・オープン時の挙動をモック
	oldOsOpenFile := compiler.OsOpenFile
	defer func() {
		compiler.OsOpenFile = oldOsOpenFile
	}()

	compiler.OsOpenFile = func(name string, flag int, perm fs.FileMode) (*os.File, error) {
		return nil, errors.New("forced error")
	}

	// 実行と標準出力・標準エラー出力のキャプチャ
	out := capturer.CaptureOutput(func() {
		// temp ディレクトリ先のスクリプト実行
		err := cpl.CompileOnly(mods, byteData, pathFileScript, "")

		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to open file")
		assert.Contains(t, err.Error(), "forced error")
	})

	assert.Empty(t, out)
}

// コンパイル済みファイルの出力先のパスが存在するもディレクトリの場合のエラーテスト。
func TestCompileOnly_not_file(t *testing.T) {
	cpl := compiler.New()

	// テスト用スクリプトのパス取得
	pathFileScript := getPathTestData(t, "hello_world.qiitan")

	// コンパイルされてないデータの読み込み
	byteData, err := cpl.ReadScript(pathFileScript)
	require.NoError(t, err)

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	// 実行と標準出力・標準エラー出力のキャプチャ
	out := capturer.CaptureOutput(func() {
		outputFile := t.TempDir() // 出力先を temp ディレクトリに

		err := cpl.CompileOnly(mods, byteData, pathFileScript, outputFile)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "is a directory")
	})

	assert.Empty(t, out)
}

// ----------------------------------------------------------------------------
//  CompileScript()
// ----------------------------------------------------------------------------

// CompileScript の正常系テスト。
func TestCompileScript(t *testing.T) {
	cpl := compiler.New()

	cpl.ResolvePath = false

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	src := []byte(`
		fmt := import("fmt")
		fmt.println("foobar")
	`)

	out := capturer.CaptureOutput(func() {
		_, err := cpl.CompileScript(mods, src, "dummy")

		require.NoError(t, err)
	})

	assert.Empty(t, out)
}

// 不正なカスタム拡張子のテスト。
// スクリプトの拡張子をカスタムするも、拡張子を設定し忘れていた場合のテスト。
func TestCompileScript_fail_import_file_ext(t *testing.T) {
	cpl := compiler.New()

	// カスタム拡張子をセット（空にセット）
	cpl.ExtensionScript = []string{}

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	src := []byte(`func() { }`)

	out := capturer.CaptureOutput(func() {
		_, err := cpl.CompileScript(mods, src, "dummy")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "missing arg: at least one argument is required")
	})

	assert.Empty(t, out)
}

// ----------------------------------------------------------------------------
//  GetBasename()
// ----------------------------------------------------------------------------

func TestGetBasename(t *testing.T) {
	cpl := compiler.New()

	for _, test := range []struct {
		path   string
		expect string
	}{
		{"foo.bar", "foo"},
		{"foo", "foo"},
		{".foo", ".foo"},
		{"/foo/bar.buzz", "bar"},
		{"/foo/bar/buzz", "buzz"},
	} {
		expect := test.expect
		actual := cpl.GetBasename(test.path)

		require.Equal(t, expect, actual)
	}
}

// ----------------------------------------------------------------------------
//  ReadScript()
// ----------------------------------------------------------------------------

// 存在しないパスの読み込みテスト。
func TestReadScript_non_existing_path(t *testing.T) {
	cpl := compiler.New()

	pathDummy := "/foo/bar/buzz"

	dat, err := cpl.ReadScript(pathDummy)

	assert.Error(t, err, "non existing path should be an error")
	assert.Empty(t, dat, "on error byte data should be empty")
}

// ----------------------------------------------------------------------------
//  RunCompiled()
// ----------------------------------------------------------------------------

// コンパイル済みのデータ実行テスト。
func TestRunCompiled(t *testing.T) {
	cpl := compiler.New()

	// コンパイルされているデータの読み込み
	inputFile := getPathTestData(t, "hello_world.qiitan")
	outputFile := filepath.Join(t.TempDir(), "hello_world.qtn")

	inputData, err := cpl.ReadScript(inputFile)
	require.NoError(t, err)

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	// コンパイルのみ
	err = cpl.CompileOnly(mods, inputData, inputFile, outputFile)
	require.NoError(t, err)
	assert.FileExists(t, outputFile)

	// コンパイル済みデータの読み込み
	byteData, err := os.ReadFile(outputFile)
	require.NoError(t, err)
	assert.NotEmpty(t, byteData)

	out := capturer.CaptureOutput(func() {
		err := cpl.RunCompiled(mods, byteData)

		require.NoError(t, err)
	})

	assert.Contains(t, out, "Hello World!")
}

// コンパイル済みではないデータ読み込みのテスト。
func TestRunCompiled_not_compiled(t *testing.T) {
	cpl := compiler.New()

	// コンパイルされてない不正なデータの読み込み
	pathFileNotCompiled := getPathTestData(t, "malformed.qiitan")

	byteData, err := cpl.ReadScript(pathFileNotCompiled)
	require.NoError(t, err)

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	out := capturer.CaptureOutput(func() {
		err := cpl.RunCompiled(mods, byteData)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "bad data")
	})

	assert.Empty(t, out)
}

// ============================================================================
//  Helper Functions
// ============================================================================

// ../../examples/ ディレクトリにあるファイルのパスを返します。ファイルが存在し
// ない場合はテストはエラーで終了します。
func getPathExample(t *testing.T, nameFile ...string) string {
	t.Helper()

	pathTmp := filepath.Join(nameFile...)
	pathTmp = filepath.Join("../../examples/", pathTmp)

	pathFile, err := filepath.Abs(pathTmp)
	require.NoError(t, err)

	return pathFile
}

// ../../testdata/ ディレクトリにあるファイルのパスを返します。ファイルが存在し
// ない場合はテストはエラーで終了します。
func getPathTestData(t *testing.T, nameFile ...string) string {
	t.Helper()

	pathTmp := filepath.Join(nameFile...)
	pathTmp = filepath.Join("../../testdata/", pathTmp)

	pathFile, err := filepath.Abs(pathTmp)
	require.NoError(t, err)

	return pathFile
}
