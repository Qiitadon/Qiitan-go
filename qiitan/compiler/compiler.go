/*
Package compiler はスクリプト・ファイルからバイトコード（中間ファイル）の生成
および tengo VM 上での実行を行うパッケージです。

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
*/
package compiler

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/KEINOS/go-utiles/util"
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/parser"
	"github.com/pkg/errors"
)

// OsOpenFile は os.OpenFile のコピーです。
//
// テスト時に OsOpenFile の挙動を変えたい場合に、代替関数を一時的に割り当ててモックに使います。
var OsOpenFile = os.OpenFile

// Compiler は、スクリプト・ファイルのバイト・コード変換と実行を行うオブジェクトを定義します。
type Compiler struct {
	ExtensionCompiled string   // バイトコード変換済みのファイルの拡張子
	ExtensionScript   []string // スクリプトの拡張子一覧
	ResolvePath       bool     // スクリプト内の相対パスを解決するか否かのフラグ（Tengo v3 でなくなる予定）
}

// ----------------------------------------------------------------------------
//  Constructor
// ----------------------------------------------------------------------------

// New は Compiler の初期化済み（デフォルト値設定済み）の新規オブジェクトを返します。
func New() *Compiler {
	return &Compiler{
		ExtensionScript:   []string{".qiitan", ".tengo"},
		ExtensionCompiled: ".qtn",
		ResolvePath:       true,
	}
}

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// CompileAndRun はスクリプトをコンパイルし、tengo VM 上で実行します。
//
// inputFile はスクリプト・ファイルのパスで、script 内の相対パスを解決するために
// 必要です。
func (c *Compiler) CompileAndRun(modules *tengo.ModuleMap, script []byte, inputFile string) (err error) {
	// スクリプトのコンパイル
	bytecode, err := c.CompileScript(modules, script, inputFile)
	if err != nil {
		return err
	}

	machine := tengo.NewVM(bytecode, nil, -1) // tengo 仮想マシンの作成

	return machine.Run() // 仮想マシン上で実行
}

// CompileOnly はスクリプトをコンパイルし、コンパイル済みバイナリ・データを
// outputFile に書き込みます。
// outputFile が空の場合は、スクリプトと同じ階層に *.qtn ファイルを出力します。
func (c *Compiler) CompileOnly(modules *tengo.ModuleMap, data []byte, inputFile, outputFile string) (err error) {
	// 出力先のパスをスクリプトと同じ階層に設定
	if outputFile == "" {
		outputFile = filepath.Join(filepath.Dir(inputFile), c.GetBasename(inputFile)+c.ExtensionCompiled)
	}

	if util.IsDir(outputFile) {
		return errors.Errorf("given file path is a directory: %v\n", outputFile)
	}

	// スクリプトをバイト・コードに変換
	bytecode, err := c.CompileScript(modules, data, inputFile)
	if err != nil {
		return errors.Wrap(err, "failed to compile script to bytecode")
	}

	out, err := OsOpenFile(outputFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to open file")
	}

	defer func() {
		if cErr := out.Close(); err == nil {
			err = cErr
		}
	}()

	fmt.Println(outputFile) // 出力先のパスを表示

	// out にバイト・データを gob 形式にエンコードして書き込み
	return bytecode.Encode(out)
}

// CompileScript はスクリプトをコンパイルし、バイトコード（中間ファイル）に変換
// されたものを返します。
//
// 引数の inputFile は script の元となったスクリプト・ファイルのパスで、スクリプト内
// の相対パスの解決に必要です。
func (c *Compiler) CompileScript(modules *tengo.ModuleMap, script []byte, inputFile string) (*tengo.Bytecode, error) {
	// コンパイル対象をファイル・セットに追加
	fileSet := parser.NewFileSet()
	sourceFile := fileSet.AddFile(filepath.Base(inputFile), -1, len(script))

	// 構造解析用のパーサー作成
	p := parser.NewParser(sourceFile, script, nil)

	// スクリプトの構造を解析し AST（抽象構文木）のユニットにパース
	fileAST, err := p.ParseFile()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse file")
	}

	// 本家 Tengo コンパイラ・オブジェクトの作成
	cpl := tengo.NewCompiler(sourceFile, nil, nil, modules, nil)
	cpl.EnableFileImport(true)

	if err = cpl.SetImportFileExt(c.ExtensionScript...); err != nil {
		return nil, err
	}

	// コマンドの --no-resolve オプション対応
	//
	// デフォルトで、外部モジュール内の相対パスは呼び出し元のスクリプトから見た
	// ときのパスに置き換えられます。スクリプトでなく現在の作業ディレクトリから
	// みた時のパスにしたい場合は、--no-resolve オプションをつけます。
	// この機能は Tengo v3 で削除される予定です。
	if !c.ResolvePath {
		// 外部スクリプトを参照する際のルート・ディレクトリをセット
		cpl.SetImportDir(filepath.Dir(inputFile))
	}

	// コンパイルの実行（AST ノードを解析します）
	if err := cpl.Compile(fileAST); err != nil {
		return nil, errors.Wrap(err, "failed to compile from file:"+fileAST.String())
	}

	bytecode := cpl.Bytecode()  // コンパイル結果をバイトコードで取得
	bytecode.RemoveDuplicates() // バイトコードから定数と重複する値を削除

	return bytecode, nil
}

// GetBasename は拡張子なしのファイル名を返します。
func (c *Compiler) GetBasename(path string) (basename string) {
	s := filepath.Base(path)

	if n := strings.LastIndexByte(s, '.'); n > 0 {
		return s[:n]
	}

	return s
}

// ReadScript は pathFile から読み込んだ byte データを返します。
func (c *Compiler) ReadScript(pathFile string) ([]byte, error) {
	inputData, err := os.ReadFile(pathFile)
	if err != nil {
		return nil, errors.Wrap(err, "error reading script file")
	}

	return inputData, nil
}

// RunCompiled はコンパイル済みのバイナリ・ファイルを読み込み tengo VM 上で実行
// します。
func (c *Compiler) RunCompiled(modules *tengo.ModuleMap, data []byte) (err error) {
	// バイトコードのデコード
	bytecode := &tengo.Bytecode{}
	if err = bytecode.Decode(bytes.NewReader(data), modules); err != nil {
		return errors.Wrap(err, "error to decode bytecode")
	}

	machine := tengo.NewVM(bytecode, nil, -1) // tengo 仮想マシンの作成

	return machine.Run() // 仮想マシン上で実行
}
