//nolint:dupl // To make it easier to follow the source, we will leave the code that seems to be duplicated intact.
package repl_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/Qiitadon/Qiitan-go/qiitan/modules"
	"github.com/Qiitadon/Qiitan-go/qiitan/repl"
	"github.com/stretchr/testify/assert"
	"github.com/zenizh/go-capturer"
)

// 対話モードの起動と正常系テスト。
func TestREPL(t *testing.T) {
	// コマンド引数のバックアップと遅延リストア
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	// テスト用スクリプトを標準入力に割り当て
	stdin := bytes.NewBufferString(`
		fmt := import("fmt")
		fmt.println("Hello" + " World!")

		a := 1
		a += 2
		fmt.println("result:", a)
		bye
	`)
	stdout := new(bytes.Buffer) // 標準出力を割り当て（キャプチャ用）

	tempREPL := repl.REPL{
		Prompt:   "> ",  // 対話モード（REPL）時のプロンプト
		QuitWord: "bye", // 対話モード（REPL）時の終了キーワード
	}

	out := capturer.CaptureStdout(func() {
		os.Args = []string{
			t.Name(),
		}

		tempREPL.Run(mods, stdin, stdout)
	})

	assert.Contains(t, stdout.String(), "Welcome to Qiitan Script Interactive Mode")
	assert.Contains(t, stdout.String(), "Nice chatting with you. Thank you")

	assert.Contains(t, out, "printf: <user-function>")
	assert.Contains(t, out, "Hello World!")
	assert.Contains(t, out, "result:3")
}

// 対話モード時の構文エラーのテスト（配列の構文エラー）。
func TestREPL_bad_element_syntax(t *testing.T) {
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	stdin := bytes.NewBufferString("[1, 2, 3,]")
	stdout := new(bytes.Buffer)

	tempREPL := repl.REPL{
		Prompt:   "> ",  // 対話モード（REPL）時のプロンプト
		QuitWord: "bye", // 対話モード（REPL）時の終了キーワード
	}

	out := capturer.CaptureStdout(func() {
		os.Args = []string{
			t.Name(),
		}

		tempREPL.Run(mods, stdin, stdout)
	})

	assert.Empty(t, out)
	assert.Contains(t, stdout.String(), "Parse Error: expected array element, found")
}

// 対話モード時の構文エラーのテスト（export 内の構文エラー）。
func TestREPL_bad_export_statement(t *testing.T) {
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	stdin := bytes.NewBufferString(`
		fmt := import("fmt")

		export { fn: func(a) { two := import("two/two"); return two.fn(a, "one")}}

		fmt.println("result:", fn("two"))
	`)
	stdout := new(bytes.Buffer)

	tempREPL := repl.REPL{
		Prompt:   "> ",  // 対話モード（REPL）時のプロンプト
		QuitWord: "bye", // 対話モード（REPL）時の終了キーワード
	}

	out := capturer.CaptureStdout(func() {
		os.Args = []string{
			t.Name(),
		}

		tempREPL.Run(mods, stdin, stdout)
	})

	assert.Contains(t, out, "print: <user-function>")
	assert.Contains(t, stdout.String(), "Compile Error: unresolved reference 'fn'")
}

// 対話モード時の構文エラーのテスト（型違いの代入）。
func TestREPL_bad_statement_usage(t *testing.T) {
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	// 代入命令（statement）の使用エラー
	stdin := bytes.NewBufferString(`
		a := 1; a += 2;
		a += "b" // int 型に string 型の代入
		fmt := import("fmt")
		fmt.println("result:", a)
	`)
	stdout := new(bytes.Buffer)

	tempREPL := repl.REPL{
		Prompt:   "> ",  // 対話モード（REPL）時のプロンプト
		QuitWord: "bye", // 対話モード（REPL）時の終了キーワード
	}

	out := capturer.CaptureOutput(func() {
		os.Args = []string{
			t.Name(),
		}

		tempREPL.Run(mods, stdin, stdout)
	})

	assert.Contains(t, stdout.String(), "Runtime Error: invalid operation: int + string")
	assert.Contains(t, out, "result:3")
}

// 対話モード時の構文エラーのテスト（二重定義）。
func TestREPL_redeclaration(t *testing.T) {
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	// 二重宣言してみる
	stdin := bytes.NewBufferString(`
		fmt := import("fmt")
		fmt := import("fmt")
		fmt.println("Hello" + " World!")
	`)
	stdout := new(bytes.Buffer)

	tempREPL := repl.REPL{
		Prompt:   "> ",  // 対話モード（REPL）時のプロンプト
		QuitWord: "bye", // 対話モード（REPL）時の終了キーワード
	}

	out := capturer.CaptureOutput(func() {
		os.Args = []string{
			t.Name(),
		}

		tempREPL.Run(mods, stdin, stdout)
	})

	assert.Contains(t, stdout.String(), "'fmt' redeclared in this block")
	assert.Contains(t, out, "Hello World!")
}
