package gomods

import (
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
)

// FibonacciModule は "fibonacci" モジュールのオブジェクト・マップです。
// ここマップに割り当てられた Go 関数が "fibonacci" モジュールの関数として利用できます。
var FibonacciModule = map[string]tengo.Object{
	"fib": &tengo.UserFunction{
		Name:  "fib",
		Value: stdlib.FuncAI64RI64(Fib),
	},
	"fibt": &tengo.UserFunction{
		Name:  "fibt",
		Value: stdlib.FuncAI64RI64(Fibt),
	},
}

// Fib はフィボナッチ数を求める関数です。
//
// qiitan スクリプトでの実行と、Go モジュールでの実行速度の違いを測るために組み込まれています。
func Fib(x int64) int64 {
	if x == 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	return Fib(x-1) + Fib(x-2)
}

// Fibt はフィボナッチ数を末尾再帰で求める関数です。
//
// qiitan スクリプトでの実行と、Go モジュールでの実行速度の違いを測るために組み込まれています。
func Fibt(x int64) int64 {
	return fibt(x, 0, 1)
}

func fibt(x, a, b int64) int64 {
	if x == 0 {
		return a
	}

	if x == 1 {
		return b
	}

	return fibt(x-1, b, a+b)
}
