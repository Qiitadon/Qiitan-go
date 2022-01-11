/*
Package repl は対話モード（REPL）時に使われるパッケージです。

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()
	// REPL インスタンスの作成
	tempREPL := repl.REPL{
		Prompt:   "> ",  // 対話モード（REPL）時のプロンプト
		QuitWord: "bye", // 対話モード（REPL）時の終了キーワード
	}
	// REPL の起動（実行）
	tempREPL.Run(mods, os.Stdin, os.Stdout)

*/
package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/parser"
)

// REPL は対話モード（REPL）でスクリプトを実行するためのオブジェクトを定義します。
type REPL struct {
	Prompt   string // 対話モード（REPL）時のプロンプト
	QuitWord string // 対話モード（REPL）時の終了キーワード
}

// Interactor は対話モード（REPL）時のループ処理で再利用されるポインタや
// オブジェクトを保持します。
// 主に InteractREPL() の引数を少なく（すっきり）させるために使われます。
type Interactor struct {
	StdIn       *bufio.Scanner
	StdOut      *io.Writer
	FileSet     *parser.SourceFileSet
	Modules     *tengo.ModuleMap
	SymbolTable *tengo.SymbolTable
	Globals     []tengo.Object
	Constants   []tengo.Object
}

// ----------------------------------------------------------------------------
//  REPL（対話モード）用ソース
// ----------------------------------------------------------------------------

// 引数で渡されたパーサーに REPL 用の println 命令を追記します。
func (r *REPL) addPrints(file *parser.File) *parser.File {
	var stmts []parser.Stmt

	for _, s := range file.Stmts {
		switch s := s.(type) {
		case *parser.ExprStmt:
			stmts = append(stmts, &parser.ExprStmt{
				Expr: &parser.CallExpr{
					Func: &parser.Ident{Name: "__repl_println__"},
					Args: []parser.Expr{s.Expr},
				},
			})
		case *parser.AssignStmt:
			stmts = append(stmts, s)

			stmts = append(stmts, &parser.ExprStmt{
				Expr: &parser.CallExpr{
					Func: &parser.Ident{
						Name: "__repl_println__",
					},
					Args: s.LHS,
				},
			})
		default:
			stmts = append(stmts, s)
		}
	}

	return &parser.File{
		InputFile: file.InputFile,
		Stmts:     stmts,
	}
}

// 引数で渡されたオブジェクトの値を標準出力に出力します。
//
// この関数は、対話モード（REPL）時、ユーザーの各アクションごとの実行結果を出力
// するのに使われます。
func gPrintln(args ...tengo.Object) (ret tengo.Object, err error) {
	printArgs := make([]interface{}, 0)

	for _, arg := range args {
		if _, isUndefined := arg.(*tengo.Undefined); isUndefined {
			printArgs = append(
				printArgs,
				"LOG: <undefined> object returned. Perhaps the statement has no return or nothing is assigned to it.")

			continue
		}

		s, _ := tengo.ToString(arg)
		printArgs = append(printArgs, "LOG: "+s)
	}

	printArgs = append(printArgs, "\n")
	_, _ = fmt.Print(printArgs...)

	return
}

// InteractREPL は対話モード（REPL）時の、ユーザーの 1 アクションごとの処理です。
// 標準入力（Interactor オブジェクトの StdIn フィールド）からコードを受け取り、
// 実行します。
func (r *REPL) InteractREPL(interactor *Interactor) (doContinue, doReturn bool) {
	// プロンプトの表示
	_, _ = fmt.Fprint(*interactor.StdOut, r.Prompt)

	if !interactor.StdIn.Scan() {
		return false, true
	}

	// ユーザ入力を取得
	line := interactor.StdIn.Text()

	// 対話モードの終了
	if strings.TrimSpace(line) == r.QuitWord {
		_, _ = fmt.Fprint(*interactor.StdOut, "Nice chatting with you. Thank you! Bye-bye~.\n")

		return false, true
	}

	// 入力を履歴に追加 & 履歴からパーサーの作成
	srcFile := interactor.FileSet.AddFile("repl", -1, len(line))
	p := parser.NewParser(srcFile, []byte(line), nil)

	file, err := p.ParseFile()
	if err != nil {
		_, _ = fmt.Fprintln(*interactor.StdOut, err.Error())

		return true, false
	}

	file = r.addPrints(file)

	c := tengo.NewCompiler(srcFile, interactor.SymbolTable, interactor.Constants, interactor.Modules, nil)
	if err := c.Compile(file); err != nil {
		_, _ = fmt.Fprintln(*interactor.StdOut, err.Error())

		return true, false
	}

	bytecode := c.Bytecode()

	machine := tengo.NewVM(bytecode, interactor.Globals, -1)
	if err := machine.Run(); err != nil {
		_, _ = fmt.Fprintln(*interactor.StdOut, err.Error())

		return true, false
	}

	interactor.Constants = bytecode.Constants

	return false, false
}

// Run は対話モードを開始（REPL で実行）します。
//
// in から受け取ったスクリプトを実行し、out に出力します。
// スクリプト実行時に使うモジュールは、予め modules に含めておく必要があります。
func (r *REPL) Run(modules *tengo.ModuleMap, in io.Reader, out io.Writer) {
	// グリーティング・メッセージの表示
	_, _ = fmt.Fprint(out, "Welcome to Qiitan Script Interactive Mode! (To exit just type `bye`)\n")

	interactor := Interactor{
		StdIn:       bufio.NewScanner(in),
		StdOut:      &out,
		FileSet:     parser.NewFileSet(),
		Modules:     modules,
		SymbolTable: tengo.NewSymbolTable(),
		Globals:     make([]tengo.Object, tengo.GlobalsSize),
		Constants:   make([]tengo.Object, 0),
	}

	// qiitan の標準モジュールを利用可能なモジュールとして割り当て
	for idx, fn := range tengo.GetAllBuiltinFunctions() {
		interactor.SymbolTable.DefineBuiltin(idx, fn.Name)
	}

	// gPrintln 関数をグローバル関数 println として埋め込み（変数の状態を適宜出力するのに利用）
	symbol := interactor.SymbolTable.Define("__repl_println__")
	interactor.Globals[symbol.Index] = &tengo.UserFunction{
		Name:  "println",
		Value: gPrintln,
	}

	for {
		if _, doReturn := r.InteractREPL(&interactor); doReturn {
			return
		}
	}
}
