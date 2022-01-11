/*
自動生成リスト

リポジトリのルートから `go generate ./...` コマンドを実行すると、下記 `go:generate <コマンド>`
の `<コマンド>` 文が実行されます。
*/
package main

// /qiitan/modules/ModulesInScript の自動生成と生成後の gofmt 実行
//go:generate go run ../gen/gen_scrmods/ -output "../qiitan/modules/ModulesInScript.go"
//go:generate go fmt ../qiitan/modules/ModulesInScript.go

// /docs/index.md の自動生成
//go:generate /bin/sh ../gen/gen_index/gen_index.sh
