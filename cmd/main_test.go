package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zenizh/go-capturer"
)

func Test_main(t *testing.T) {
	// backup and defer restore
	oldOsArgs := os.Args
	oldVersionFlag := versionFlag

	defer func() {
		os.Args = oldOsArgs
		versionFlag = oldVersionFlag
	}()

	// mock args
	os.Args = []string{
		t.Name(),
	}

	out := capturer.CaptureStdout(func() {
		main()
	})

	assert.Equal(t, "Hello, Qiitan!\n", out)
}

func Test_version(t *testing.T) {
	// backup and defer restore
	oldOsArgs := os.Args
	oldVersionFlag := versionFlag

	defer func() {
		os.Args = oldOsArgs
		versionFlag = oldVersionFlag
	}()

	// mock args
	os.Args = []string{
		t.Name(),
		"--version",
	}

	out := capturer.CaptureStdout(func() {
		main()
	})

	assert.Contains(t, out, "cmd")
	assert.Contains(t, out, "(devel)")
}

func Test_version_defined(t *testing.T) {
	// backup and defer restore
	oldOsArgs := os.Args
	oldVersionFlag := versionFlag
	oldVersion := version

	defer func() {
		os.Args = oldOsArgs
		versionFlag = oldVersionFlag
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

	assert.Contains(t, out, "cmd")
	assert.Contains(t, out, "0.0.0")
}
