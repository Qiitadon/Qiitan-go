package modules

import (
	"github.com/Qithub-BOT/Qiitan-go/qiitan/modules/coremods"
	"github.com/d5/tengo/v2"
)

// GetModuleNamesAll はデフォルト（標準）のモジュール名一覧を返します。
func GetModuleNamesAll() []string {
	_, names := GetModuleAll()

	return names
}

// GetModuleMap は引数で指定されたモジュール名を含んだ tengo.ModuleMap のオブ
// ジェクトを返します。
func GetModuleMap(names ...string) *tengo.ModuleMap {
	result := tengo.NewModuleMap()
	mods, _ := GetModuleAll()

	for _, name := range names {
		// Go 言語によるモジュールを追加
		if mod := mods.GetBuiltinModule(name); mod != nil {
			result.AddBuiltinModule(name, mod.Attrs)
		}

		// Tengo 言語によるモジュールを追加
		if mod := mods.GetSourceModule(name); mod != nil {
			result.AddSourceModule(name, mod.Src)
		}
	}

	return result
}

// GetModuleMapAll はデフォルト（標準）のモジュールを含んだモジュール・マップの
// オブジェクトを返します。
func GetModuleMapAll() *tengo.ModuleMap {
	result, _ := GetModuleAll()

	return result
}

// GetModuleAll はデフォルト（標準）のモジュールを含んだモジュール・マップの
// オブジェクトおよびモジュール名の一覧を返します。
func GetModuleAll() (*tengo.ModuleMap, []string) {
	mods := tengo.NewModuleMap()
	names := []string{}

	// Go 言語で書かれたモジュール一覧を追加
	for _, modsInGo := range []map[string]map[string]tengo.Object{
		coremods.InGo(),
		ModulesInGo,
	} {
		for name, mod := range modsInGo {
			names = append(names, name)
			mods.AddBuiltinModule(name, mod)
		}
	}

	// Tengo 言語で書かれたモジュール一覧を追加
	for _, modsInTengo := range []map[string]string{
		coremods.InTengo(),
		ModulesInScript,
	} {
		for name, script := range modsInTengo {
			names = append(names, name)
			mods.AddSourceModule(name, []byte(script))
		}
	}

	return mods, names
}
