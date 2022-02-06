package gomods_test

import (
	"context"
	"testing"

	"github.com/Qiitadon/Qiitan-go/qiitan/modules"
	"github.com/d5/tengo/v2"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  Helper Function
// ----------------------------------------------------------------------------

// RunScript はスクリプトを実行して、実行後の各オブジェクトの状態を返します。
func RunScript(t *testing.T, script string) *tengo.Compiled {
	t.Helper()

	s := tengo.NewScript([]byte(script))
	s.SetImports(modules.GetModuleMapAll())

	// run the script
	compiled, err := s.RunContext(context.Background())
	require.NoError(t, err)

	return compiled
}
