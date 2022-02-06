/*
Package main は、ドキュメント内にある qiitan スクリプトをコードブロックごとにテスト実行します。

1. 引数で渡されたディレクトリ以下にある Markdown ドキュメントの一覧を取得し、各ドキュメントごとに走査します。
2. コードブロックの言語シンタックスが `go qiitan` の場合に実行され、ステータス `0`（ゼロ）で終了した場合にパスします。
3. コードブロック内に `Output:` コメント行があった場合、続くコメント行と実際の標準出力の内容を比較し、同じであった場合にパスします。

*/
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/KEINOS/go-utiles/util"
	"github.com/pkg/errors"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func main() {
	flag.Parse()

	listDirSearch := flag.Args()

	util.ExitOnErr(Run(listDirSearch))
}

// GetListPathFile returns a file path list of the markdown files to test.
func GetListPathFile(listDirSearch []string) ([]string, error) {
	lenDir := len(listDirSearch)
	if lenDir == 0 {
		return []string{}, errors.New("the directory to search is not specified")
	}

	pathFileHolder := make([]string, lenDir)

	for _, pathDir := range listDirSearch {
		listFile, err := SearchDir(pathDir)
		if err != nil {
			return []string{}, errors.Wrap(err, "failed to search directory")
		}

		pathFileHolder = append(pathFileHolder, listFile...)
	}

	return pathFileHolder, nil
}

// GetListScript returns a list of markdown code blocks found in pathFile.
func GetListScript(pathFile string) (result []string, err error) {
	fileByte, err := os.ReadFile(pathFile)
	if err != nil {
		return []string{}, errors.Wrap(err, "failed to read file")
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithBlockParsers(),
			parser.WithInlineParsers(),
		),
	)

	markdownParser := markdown.Parser()
	textReader := text.NewReader(fileByte)
	astNodes := markdownParser.Parse(textReader)
	level := int(0)

	ast.Walk(astNodes, func(currNode ast.Node, entering bool) (ast.WalkStatus, error) {
		var (
			err           error
			typeDesc      string
			enteringStats string
		)

		status := ast.WalkStatus(ast.WalkContinue)
		nodeKind := currNode.Kind().String()

		if entering {
			enteringStats = "IN :"
			level++
		} else {
			enteringStats = "OUT:"
			level--
		}

		//textFound := util.FmtStructPretty(currNode)
		typeDesc = fmt.Sprintf("%T", currNode)
		textFound := util.FmtStructPretty(currNode)

		switch currNode.(type) {
		case *ast.Document:
			typeDesc = "it is a doc"
		case *ast.Text:
			typeDesc = "it is a text"
		case *ast.Paragraph:
			typeDesc = "it is a Paragraph"
		case *ast.FencedCodeBlock:
			typeDesc = "it is a FencedCodeBlock"
		case *ast.CodeBlock:
			typeDesc = "it is a inline"
		default:
			typeDesc = fmt.Sprintf("Unused Type: %T", currNode)
		}

		msg := fmt.Sprintf(
			"%v:%v Kind: %v, Desc: %v, Text: %v",
			level,
			enteringStats,
			nodeKind,
			typeDesc,
			textFound,
		)

		result = append(result, msg)

		return status, err
	})

	return result, nil
}

// Run is the actual main function of this package.
func Run(listDirSearch []string) error {
	// Get lists of markdown files
	ListPathFile, err := GetListPathFile(listDirSearch)
	if err != nil {
		return err
	}

	sort.Strings(ListPathFile)

	fmt.Println(ListPathFile)

	return nil
}

func RunScript(script string) error {
	return nil
}

func SearchDir(pathDir string) ([]string, error) {
	pathDir, err := filepath.Abs(pathDir)

	if err != nil || !util.IsDir(pathDir) {
		return []string{}, errors.New("dir not exists: " + pathDir)
	}

	return []string{pathDir}, nil
}
