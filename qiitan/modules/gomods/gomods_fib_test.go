package gomods_test

import (
	"testing"

	"github.com/Qithub-BOT/Qiitan-go/qiitan/modules/gomods"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func Test_fib_fib(t *testing.T) {
	script := `
		fib := import("fibonacci")
		result := fib.fib(35)
	`
	compiled := RunScript(t, script)

	// 実行結果の取り出し
	actual := compiled.Get("result")
	expect := "9227465"

	assert.Equal(t, expect, actual.String())
}

func Test_fib_fibt(t *testing.T) {
	script := `
		fib := import("fibonacci")
		result := fib.fibt(35)
	`
	compiled := RunScript(t, script)

	// 実行結果の取り出し
	actual := compiled.Get("result")
	expect := "9227465"

	assert.Equal(t, expect, actual.String())
}

func TestFib(t *testing.T) {
	expect := int64(9227465)
	actual := gomods.Fib(35)

	assert.Equal(t, expect, actual)
}

func TestFibt(t *testing.T) {
	expect := int64(9227465)
	actual := gomods.Fibt(35)

	assert.Equal(t, expect, actual)
}

func TestFibt_zero(t *testing.T) {
	expect := int64(0)
	actual := gomods.Fibt(0)

	assert.Equal(t, expect, actual)
}
