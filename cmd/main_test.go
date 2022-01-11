package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenizh/go-capturer"
)

func Test_main(t *testing.T) {
	out := capturer.CaptureStdout(func() {
		main()
	})

	assert.Equal(t, out, "Hello, Qiitan!\n")
}

func Test_version(t *testing.T) {
	// backup and defer restore
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	// mock args
	os.Args = []string{
		t.Name(),
		"--version",
	}

	out := capturer.CaptureStdout(func() {
		main()
	})

	assert.Contains(t, out, "cmd.test (devel)")
}

func Test_version_defined(t *testing.T) {
	// backup and defer restore
	oldOsArgs := os.Args
	defer func() {
		os.Args = oldOsArgs
	}()

	oldVersion := version
	defer func() {
		version = oldVersion
	}()

	// mock Args and version
	os.Args = []string{
		t.Name(),
		"--version",
	}

	version = "0.0.0"

	out := capturer.CaptureStdout(func() {
		main()
	})

	assert.Contains(t, out, "cmd.test 0.0.0")
}
