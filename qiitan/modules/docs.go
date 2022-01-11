/*
Package modules は qiitan インタプリタに標準で同梱されるモジュールを定義します。

Tengo 言語の標準モジュール（tengo.stdlib）に加え、Qiitan スクリプト用に独自に拡
張したモジュールが定義されています。

About module

Qiitan スクリプトにおけるモジュールとは Qiitan スクリプトの言語を拡張するライブ
ラリ（関数群）です。

標準で同梱されているライブラリは Qiitan スクリプト内の `import()` で指定する
だけでスクリプト内で使用できます。

標準モジュールは、 Go 言語および Tengo 言語で記述できます.

About tengo (source) module

Tengo 言語で書かれたモジュールのファイル名は "scrmod_" ではじまり、".tengo" で
終わる必要があります。

    scrmod_<パッケージ名>.tengo
*/
package modules

//go:generate go run ../../gen/gen_scrmods/ -output "../../qiitan/modules/ModulesInScript.go"
//go:generate go fmt ../../qiitan/modules/ModulesInScript.go
