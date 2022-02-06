package coremods_test

import (
	"testing"

	"github.com/Qiitadon/Qiitan-go/qiitan/modules/coremods"
	"github.com/stretchr/testify/assert"
)

func TestInGo(t *testing.T) {
	tmpMods := coremods.InGo()

	for _, test := range []struct {
		nameMod    string
		nameMethod string
	}{
		{"base64", "encode"},
		{"fmt", "printf"},
		{"hex", "encode"},
		{"json", "encode"},
		{"math", "pi"},
		{"os", "chown"},
		{"rand", "rand"},
		{"text", "compare"},
		{"times", "sleep"},
	} {
		tmpMod := tmpMods[test.nameMod]
		tmpFunc := tmpMod[test.nameMethod]

		assert.NotNil(t, tmpMod,
			"stdlib shuould contain module: %v", test.nameMod)
		assert.NotNil(t, tmpFunc,
			"module %v should contain function: %v", test.nameMod, test.nameMethod)
	}
}

func TestInTengo(t *testing.T) {
	tmpMods := coremods.InTengo()

	for _, test := range []struct {
		nameMod    string
		nameMethod string
	}{
		{"enum", "is_enumerable"},
	} {
		tmpSource := tmpMods[test.nameMod]

		assert.NotNil(t, tmpSource,
			"stdlib shuould contain module: %v", test.nameMod)
		assert.Contains(t, tmpSource, test.nameMethod,
			"module %v should contain function: %v", test.nameMod, test.nameMethod)
	}
}
