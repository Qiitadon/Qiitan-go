package gomods_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Tests
// ----------------------------------------------------------------------------

func Test_hash_md5(t *testing.T) {
	script := `
		hash := import("hash")
		result := hash.md5("beef")
	`
	compiled := RunScript(t, script)

	// 実行結果の取り出し
	actual := compiled.Get("result")
	expect := "34902903de8d4fee8e6afe868982f0dd"

	assert.Equal(t, expect, actual.String())
}

func Test_hash_sha1(t *testing.T) {
	script := `
		hash := import("hash")
		result := hash.sha1("beef")
	`
	compiled := RunScript(t, script)

	// 実行結果の取り出し
	actual := compiled.Get("result")
	expect := "1b57bbec9e3484d6678c162410d0e4da9f8af6f0"

	assert.Equal(t, expect, actual.String())
}

func Test_hash_sha224(t *testing.T) {
	script := `
		hash := import("hash")
		result := hash.sha224("beef")
	`
	compiled := RunScript(t, script)

	// 実行結果の取り出し
	actual := compiled.Get("result")
	expect := "46645a3aad2c42a3668f1fa7ff4e703a3c6f181827036b65d8c72af4"

	assert.Equal(t, expect, actual.String())
}

func Test_hash_sha256(t *testing.T) {
	script := `
		hash := import("hash")
		result := hash.sha256("beef")
	`
	compiled := RunScript(t, script)

	// 実行結果の取り出し
	actual := compiled.Get("result")
	expect := "aa415c4e8890cf0fec7826aec962ffbcc04534faefd2b3266c54f690d40d6e82"

	assert.Equal(t, expect, actual.String())
}

func Test_hash_sha384(t *testing.T) {
	script := `
		hash := import("hash")
		result := hash.sha384("beef")
	`
	compiled := RunScript(t, script)

	// 実行結果の取り出し
	actual := compiled.Get("result")
	expect := "6bd411a8c3a1177b4a51ce612c17b178e1da7a059401b1af4383a4a33872e5ee" +
		"7b10038ced7e1414c32e95b61e6714f9"

	assert.Equal(t, expect, actual.String())
}

func Test_hash_sha512(t *testing.T) {
	script := `
		hash := import("hash")
		result := hash.sha512("beef")
	`
	compiled := RunScript(t, script)

	// 実行結果の取り出し
	actual := compiled.Get("result")
	expect := "8cd8bb0cef938ef9cd054c2c2cb965e83310ab5c197cb5fc8f35892a44c1a028" +
		"bac9e1bcd6248580fa2739cc96074885ea3ee116ef35c2d8f6124270aeff50b7"

	assert.Equal(t, expect, actual.String())
}
