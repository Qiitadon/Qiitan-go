package modules

import (
	"github.com/Qiitadon/Qiitan-go/qiitan/modules/gomods"
	"github.com/d5/tengo/v2"
)

// ModulesInGo は Go 言語で書かれた、キーたん語の標準モジュール一覧を保持します。
//
// "./gomods/" 下にモジュールを設置した場合は、ここにも追記する必要があります。
var ModulesInGo = map[string]map[string]tengo.Object{
	"hash":      gomods.HashModule,
	"fibonacci": gomods.FibonacciModule,
}
