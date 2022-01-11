package modules_test

import (
	"context"
	"testing"

	"github.com/Qithub-BOT/Qiitan-go/qiitan/modules"
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/require"
	"github.com/stretchr/testify/assert"
)

func Test_srcmod_hello(t *testing.T) {
	src := `
	hello := import("hello")

	result1 := hello.world("foo", "bar")
	result2 := hello.qiitan()
`
	s := tengo.NewScript([]byte(src))

	err := s.Add("description", "undefined")
	require.NoError(t, err)

	s.SetImports(modules.GetModuleMapAll())

	// run the script
	compiled, err := s.RunContext(context.Background())
	require.NoError(t, err)

	{
		// retrieve values
		actual := compiled.Get("result1")
		expect := "Hello, foo bar!"

		assert.Equal(t, expect, actual.String())
	}
	{
		actual := compiled.Get("result2")
		expect := "Hi! How are you?"

		assert.Equal(t, expect, actual.String())
	}
}
