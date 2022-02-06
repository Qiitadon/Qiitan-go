package main

import (
	"os"
	"testing"

	"github.com/KEINOS/go-utiles/util"
	"github.com/d5/tengo/v2/require"
	"github.com/stretchr/testify/assert"
	"github.com/zenizh/go-capturer"
)

func Test_main_no_arg(t *testing.T) {
	oldOsArg := os.Args
	defer func() {
		os.Args = oldOsArg
	}()

	oldOsExit := util.OsExit
	defer func() {
		util.OsExit = oldOsExit
	}()

	// Mock args
	os.Args = []string{
		t.Name(), // As bin name
	}

	// Mock os.Exit
	var capturedStatus int

	util.OsExit = func(code int) {
		capturedStatus = code
	}

	// Test
	out := capturer.CaptureStderr(func() {
		main()
	})

	// Assertion
	assert.Equal(t, 1, capturedStatus, "it should end with status 1 if no arguments was set")
	assert.Contains(t, out, "the directory to search is not specified")
}

func TestGetListScript(t *testing.T) {
	pathFile := "./README.md"

	listScript, err := GetListScript(pathFile)
	require.NoError(t, err)

	t.Log(util.FmtStructPretty(listScript))

	t.Fail()
}

func TestRunScript(t *testing.T) {
	script := util.HereDoc(`
		fmt := import("fmt")
		fmt.println("Hello, world! from RunScript")
	`)

	err := RunScript(script)
	require.NoError(t, err)
}
