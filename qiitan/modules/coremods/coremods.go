/*
Package coremods は「Tengo 言語の標準モジュール」`stdlib.BuiltinModules` および
`stdlib.SourceModules`を、キーたん語のコア・モジュールとして組み込みます。

これらのモジュールは Tengo スクリプトと互換性を持ちます。
*/
package coremods

import (
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
)

// InGo は Tengo のビルトイン・モジュール（標準モジュール）のうち、Go
// 言語で書かれたモジュールを返します。
func InGo() map[string]map[string]tengo.Object {
	return stdlib.BuiltinModules
}

// InTengo は Tengo のビルトイン・モジュール（標準モジュール）のうち、Tengo
// 言語で書かれたモジュールを返します。
//
// Tengo ではスクリプトで書かれたモジュールを Source モジュールと呼んでいます
// が、キーたん語では Script モジュールと呼びます。
func InTengo() map[string]string {
	return stdlib.SourceModules
}
