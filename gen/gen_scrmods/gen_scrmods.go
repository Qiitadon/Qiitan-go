package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/KEINOS/go-utiles/util"
	"github.com/pkg/errors"
)

const nameFileOutDefault = "ModulesInScript.go" // 出力するファイル名のデフォルト

var (
	tengoModFileRE = regexp.MustCompile(`^scrmod_(\w+).tengo$`)
	nameFileOut    = flag.String("output", "", "output file name")
)

func main() {
	flag.Parse()

	pathDirOut := getPathDirOut()
	pathDirIn := filepath.Join(pathDirOut, "scrmods")

	err := GenerateScriptModules(pathDirIn, pathDirOut)
	util.ExitOnErr(err)
}

// GenerateScriptModules は /modules/ModulesInScript.go のジェネレーターです。
//
// pathDirIn 下にある *.tengo モジュール群を読み込み、 ModulesInScript.go を
// pathDirOut に作成します。
func GenerateScriptModules(pathDirIn, pathDirOut string) error {
	modules := make(map[string]string)

	// 同梱する Tengo モジュール・ファイル群の読み込み
	files, err := os.ReadDir(pathDirIn)
	if err != nil {
		return err
	}

	for _, file := range files {
		m := tengoModFileRE.FindStringSubmatch(file.Name())
		if m != nil {
			modName := m[1]
			pathFilescr := filepath.Join(pathDirIn, file.Name())

			scr, err := os.ReadFile(pathFilescr)
			if err != nil {
				return errors.Errorf("file '%s' read error: %s",
					file.Name(), err.Error())
			}

			modules[modName] = string(scr)
		}
	}
	// ----------------------------------------------------------------------------
	var out bytes.Buffer

	// ヘッダの追加
	out.WriteString(`// Code generated using genscrmods.go; DO NOT EDIT.
//
// genscrmods.go により自動生成されたコードです。編集しないでください。
// モジュールを追加したら go generete ./... を実行して更新してください。
package modules

// ModulesInScript は Tengo 言語で書かれた、qiitan スクリプト独自のモジュールの
// 一覧を保持します。
var ModulesInScript = map[string]string{` + "\n")

	// 本体コードの追加
	for modName, modscr := range modules {
		out.WriteString("\t\"" + modName + "\": " +
			strconv.Quote(modscr) + ",\n")
	}

	out.WriteString("}\n")

	// 書き込み
	pathFileTarget := filepath.Join(pathDirOut, getNameOut())

	//nolint:gosec // 基本的に変更はしないものの、0600 でなく、他のソースとおな
	//じ 0644 にあえて指定
	return os.WriteFile(pathFileTarget, out.Bytes(), 0o644)
}

// 出力するファイルのファイル名を返します。
//
// -output フラグ・オプションで指定があった場合は、その値を返します。
// 指定がない、もしくは空の場合は、デフォルトのファイル名を返します。
func getNameOut() string {
	nameFile := nameFileOutDefault

	if *nameFileOut != "" {
		nameFile = *nameFileOut
	}

	return filepath.Base(nameFile)
}

// 出力先のディレクトリを絶対パスで返します。
func getPathDirOut() string {
	pathFile := filepath.Join(".", nameFileOutDefault)
	nameFile := *nameFileOut

	if nameFile != "" {
		pathFile = nameFile
	}

	pathDir, err := filepath.Abs(filepath.Dir(pathFile))
	util.ExitOnErr(err)

	fmt.Println("PATH:", pathDir)
	fmt.Println("ORIGIN:", nameFile)

	return pathDir
}
