package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/KEINOS/go-utiles/util"
	"github.com/Qiitadon/Qiitan-go/qiitan/compiler"
	"github.com/Qiitadon/Qiitan-go/qiitan/modules"
	"github.com/Qiitadon/Qiitan-go/qiitan/repl"
)

const (
	replPrompt      = "> "   // 対話モード（REPL）時のプロンプト
	replQuitWord    = "bye"  // 対話モード（REPL）時の終了キーワード
	compiledFileExt = ".qtn" // 中間ファイルの拡張子
)

var (
	outputFile    string
	sourceFileExt = []string{".qiitan", ".tengo"} // Qiitan スクリプトの拡張子
	flags         *flag.FlagSet
	compileOnly   bool
	noResolvePath bool
	showHelp      bool
	showVersion   bool
	version       = "dev"
)

// ----------------------------------------------------------------------------
//  Main
// ----------------------------------------------------------------------------

func main() {
	util.ExitOnErr(Run())
}

// ----------------------------------------------------------------------------
//  Functions
// ----------------------------------------------------------------------------
// ヘルプを表示します。
func doHelp() {
	msg := util.HereDoc(`
		Usage:
		  qiitan [flags] [input-file]

		Flags:
		  -o output-file  バイトコード（中間ファイル）を出力します
		  -version        qiitan インタプリタのバージョン情報を表示します

		Input Files:
		  qiitan スクリプトのファイル・パスを指定します。フラグの指定がない場合は、コン
		  パイル（中間言語変換）と実行を行います。
		  スクリプトの拡張子は ".qiitan" および ".tengo" である必要があります。ファイル
		  の指定がない場合は、対話モード（REPL）で起動します。

		REPL Mode:
		  ファイルの指定がない場合、REPL モードで起動し対話シェルを通してスクリプトを直
		  接入力して実行できます。
		  対話モードを終了するには "bye" をタイプするか ctl-c で強制終了させます。

		Examples:
		  対話モード（REPL）で qiitan インタプリタを起動します。

		    $ qiitan

		  スクリプト（myapp.qiitan）をコンパイルし実行します。拡張子は ".qiitan" および
		  ".tengo" である必要があります。

		    $ qiitan ./myapp.qiitan

		  スクリプトをコンパイルし、バイトコードの中間ファイルを出力します。（myapp.qtn）
		  外部モジュール（外部スクリプト）なども、この中間ファイルに含まれ、単体のファイ
		  ルとして出力されます。

		    $ qiitan -o myapp ./myapp.qiitan

		  バイトコード（中間ファイル）を実行します。

		    $ qiitan myapp.qtn
		`)

	fmt.Println(msg)
}

// コマンドのフラグ／オプションの初期化をします。
func initFlags() error {
	flags = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	flags.BoolVar(&showHelp, "help", false, "ヘルプを表示します。")
	flags.BoolVar(&compileOnly, "compile", false, "中間ファイルを実行せずにファイル出力します。（*.qiitan --> *.qtn）")
	flags.BoolVar(&noResolvePath, "no-resolve", false, "スクリプト内の相対パスをスクリプトを中心に展開せず、カレントのまま解釈します。")
	flags.BoolVar(&showVersion, "version", false, "バージョン情報を表示します。")
	flags.StringVar(&outputFile, "o", "", "中間ファイルを実行せずに指定したファイルのパスに出力します。")

	return flags.Parse(os.Args[1:])
}

// PreRun はスクリプトの実行以外の前処理を行います。
// 主にヘルプ表示やバージョン情報などを行います。
func PreRun() (doExit bool, err error) {
	if err := initFlags(); err != nil {
		return true, err
	}

	if showHelp {
		doHelp()

		return true, nil
	}

	if showVersion {
		fmt.Println(version)

		return true, nil
	}

	return false, nil
}

// Run は Qiitan の本体です。対話モードで実行するか、スクリプト・ファイルを読み
// 込んで実行するかの切り分けを行い、実行します。
func Run() error {
	// ヘルプやバージョン情報表示などのスクリプト実行以外の事前処理
	if doExit, err := PreRun(); doExit {
		return err
	}

	// 引数なしの場合は対話モード（REPL）で実行
	inputFile := flags.Arg(0)
	if inputFile == "" {
		// 対話モード（REPL）用のオブジェクト作成
		objREPL := repl.REPL{
			Prompt:   replPrompt,   // 対話時のプロンプト
			QuitWord: replQuitWord, // REPL モードの終了キーワード
		}

		mods := modules.GetModuleMapAll() // 標準モジュールの読み込み

		objREPL.Run(mods, os.Stdin, os.Stdout) // REPL モードで実行

		return nil
	}

	return RunScript(inputFile)
}

// RunScript は inputFile のスクリプト・ファイルを読み込み実行します。
func RunScript(inputFile string) error {
	cpl := compiler.New()

	cpl.ExtensionScript = sourceFileExt
	cpl.ExtensionCompiled = compiledFileExt

	inputData, err := cpl.ReadScript(inputFile) // ファイルの読み込み
	util.ExitOnErr(err)

	inputFile, err = filepath.Abs(inputFile) // ファイルの絶対パス取得
	util.ExitOnErr(err)

	// Shebang 対応
	if len(inputData) > 1 && string(inputData[:2]) == "#!" {
		copy(inputData, "//")
	}

	// 標準モジュールの読み込み
	mods := modules.GetModuleMapAll()

	// コンパイルのみ（-compile オプション指定時）
	if compileOnly || outputFile != "" {
		return cpl.CompileOnly(mods, inputData, inputFile, outputFile)
	}

	// コンパイルと実行（拡張子付きの場合）
	for _, ext := range sourceFileExt {
		if filepath.Ext(inputFile) == ext {
			return cpl.CompileAndRun(mods, inputData, inputFile)
		}
	}

	// コンパイル済みコードの実行
	return cpl.RunCompiled(mods, inputData)
}
