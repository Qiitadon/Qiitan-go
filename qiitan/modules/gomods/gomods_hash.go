package gomods

import (
	//nolint:gosec // weak cryptographic use on purpose
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"

	//nolint:gosec // weak cryptographic use on purpose
	"crypto/sha1"
	"fmt"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
)

// HashModule は "hash" モジュールのオブジェクト・マップです。
// ここマップに割り当てられた Go 関数が "hash" モジュールの関数として利用できます。
var HashModule = map[string]tengo.Object{
	"md5": &tengo.UserFunction{
		Name:  "md5",
		Value: stdlib.FuncASRS(HashMD5),
	},
	"sha1": &tengo.UserFunction{
		Name:  "sha1",
		Value: stdlib.FuncASRS(HashSHA1),
	},
	"sha224": &tengo.UserFunction{
		Name:  "sha224",
		Value: stdlib.FuncASRS(HashSHA224),
	},
	"sha256": &tengo.UserFunction{
		Name:  "sha256",
		Value: stdlib.FuncASRS(HashSHA256),
	},
	"sha384": &tengo.UserFunction{
		Name:  "sha384",
		Value: stdlib.FuncASRS(HashSHA384),
	},
	"sha512": &tengo.UserFunction{
		Name:  "sha512",
		Value: stdlib.FuncASRS(HashSHA512),
	},
}

// HashMD5 は MD5 ハッシュ関数の実装です。
func HashMD5(input string) string {
	//nolint:gosec // weak cryptographic use on purpose
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}

// HashSHA1 は SHA-1 ハッシュ関数の実装です。
func HashSHA1(input string) string {
	//nolint:gosec // weak cryptographic use on purpose
	return fmt.Sprintf("%x", sha1.Sum([]byte(input)))
}

// HashSHA224 は SHA-2 の SHA-224 ハッシュ関数の実装です。
func HashSHA224(input string) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(input)))
}

// HashSHA256 は SHA-2 の SHA-256 ハッシュ関数の実装です。
func HashSHA256(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}

// HashSHA384 は SHA-2 の SHA-384 ハッシュ関数の実装です。
func HashSHA384(input string) string {
	return fmt.Sprintf("%x", sha512.Sum384([]byte(input)))
}

// HashSHA512 は SHA-2 の SHA-512 ハッシュ関数の実装です。
func HashSHA512(input string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(input)))
}
