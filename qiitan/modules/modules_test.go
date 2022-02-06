package modules_test

// List of contributors sample
//     https://api.github.com/repos/d5/tengo/contributors

import (
	"testing"

	"github.com/Qiitadon/Qiitan-go/qiitan/modules"
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/require"
	"github.com/stretchr/testify/assert"
)

type (
	ARR  = []interface{}
	MAP  = map[string]interface{}
	IARR []interface{}
	IMAP map[string]interface{}
)

func TestGetModuleNamesAll(t *testing.T) {
	modNames := modules.GetModuleNamesAll()

	for _, expectName := range []string{
		// Tengo Standard Module Library
		"os", "text", "rand", "math", "times",
		"fmt", "json", "base64", "hex", "enum",
		// Extended Modules (Qiitan Standard Modules)
		"hello",
	} {
		assert.Contains(t, modNames, expectName)
	}
}

func TestGetModuleMapAll(t *testing.T) {
	mods := modules.GetModuleMapAll()
	mod := mods.GetBuiltinModule("fmt")
	obj := mod.Attrs["printf"]

	assert.IsType(t, new(tengo.UserFunction), obj)
}

func TestGetModuleMap(t *testing.T) {
	// Request module map with "fmt" and "enum"
	result := modules.GetModuleMap("fmt", "enum")

	require.NotNil(t, result)

	assert.NotNil(t, result.Get("fmt"), "it should contain the requested module")
	assert.NotNil(t, result.Get("enum"), "it should contain the requested module")
	assert.Nil(t, result.Get("os"), "it should not contain the non requested module")
}

// ----------------------------------------------------------------------------
//  Helper Functions
// ----------------------------------------------------------------------------
/*
func useModule(t *testing.T, moduleName string) callres {
	t.Helper()

	mod := modules.GetModuleMapAll()
	if mod == nil {
		return callres{t: t, e: fmt.Errorf("module not found: %s", moduleName)}
	}

	return callres{t: t, o: mod}
}

type callres struct {
	t *testing.T
	o interface{}
	e error
}

func (c callres) call(funcName string, args ...interface{}) callres {
	if c.e != nil {
		return c
	}

	oargs := []tengo.Object{}

	for _, v := range args {
		oargs = append(oargs, object(v))
	}

	switch o := c.o.(type) {
	case *tengo.BuiltinModule:
		m, ok := o.Attrs[funcName]
		if !ok {
			return callres{t: c.t, e: fmt.Errorf(
				"function not found: %s", funcName)}
		}

		f, ok := m.(*tengo.UserFunction)
		if !ok {
			return callres{t: c.t, e: fmt.Errorf("non-callable: %s", funcName)}
		}

		res, err := f.Value(oargs...)

		return callres{t: c.t, o: res, e: err}
	case *tengo.UserFunction:
		res, err := o.Value(oargs...)

		return callres{t: c.t, o: res, e: err}
	case *tengo.ImmutableMap:
		m, ok := o.Value[funcName]
		if !ok {
			return callres{t: c.t, e: fmt.Errorf("function not found: %s", funcName)}
		}

		f, ok := m.(*tengo.UserFunction)
		if !ok {
			return callres{t: c.t, e: fmt.Errorf("non-callable: %s", funcName)}
		}

		res, err := f.Value(oargs...)

		return callres{t: c.t, o: res, e: err}
	default:
		panic(fmt.Errorf("unexpected object: %v (%T)", o, o))
	}
}

func (c callres) expect(expected interface{}, msgAndArgs ...interface{}) {
	require.NoError(c.t, c.e, msgAndArgs...)
	require.Equal(c.t, object(expected), c.o, msgAndArgs...)
}

func (c callres) expectError() {
	require.Error(c.t, c.e)
}

func object(v interface{}) tengo.Object {
	switch v := v.(type) {
	case tengo.Object:
		return v
	case string:
		return &tengo.String{Value: v}
	case int64:
		return &tengo.Int{Value: v}
	case int: // for convenience
		return &tengo.Int{Value: int64(v)}
	case bool:
		if v {
			return tengo.TrueValue
		}

		return tengo.FalseValue
	case rune:
		return &tengo.Char{Value: v}
	case byte: // for convenience
		return &tengo.Char{Value: rune(v)}
	case float64:
		return &tengo.Float{Value: v}
	case []byte:
		return &tengo.Bytes{Value: v}
	case MAP:
		objs := make(map[string]tengo.Object)

		for k, v := range v {
			objs[k] = object(v)
		}

		return &tengo.Map{Value: objs}
	case ARR:
		var objs []tengo.Object

		for _, e := range v {
			objs = append(objs, object(e))
		}

		return &tengo.Array{Value: objs}
	case IMAP:
		objs := make(map[string]tengo.Object)

		for k, v := range v {
			objs[k] = object(v)
		}

		return &tengo.ImmutableMap{Value: objs}
	case IARR:
		var objs []tengo.Object

		for _, e := range v {
			objs = append(objs, object(e))
		}

		return &tengo.ImmutableArray{Value: objs}
	case time.Time:
		return &tengo.Time{Value: v}
	case []int:
		var objs []tengo.Object

		for _, e := range v {
			objs = append(objs, &tengo.Int{Value: int64(e)})
		}

		return &tengo.Array{Value: objs}
	}

	panic(fmt.Errorf("unknown type: %T", v))
}

func expect(t *testing.T, input string, expected interface{}) {
	t.Helper()

	s := tengo.NewScript([]byte(input))
	s.SetImports(modules.GetModuleMapAll())

	c, err := s.Run()

	require.NoError(t, err)
	require.NotNil(t, c)

	v := c.Get("out")

	require.NotNil(t, v)
	require.Equal(t, expected, v.Value())
}
*/
