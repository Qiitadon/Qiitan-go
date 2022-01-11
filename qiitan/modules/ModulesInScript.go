// Code generated using genscrmods.go; DO NOT EDIT.
//
// genscrmods.go により自動生成されたコードです。編集しないでください。
// モジュールを追加したら go generete ./... を実行して更新してください。
package modules

// ModulesInScript は Tengo 言語で書かれた、qiitan スクリプト独自のモジュールの
// 一覧を保持します。
var ModulesInScript = map[string]string{
	"hello": "// Module hello is a sample module written in Tengo language.\n\n// join is a private function that joins the args glued by a space.\njoin := func(...args) {\n    text := import(\"text\")\n\n    return text.join(args, \" \")\n}\n\n// qiitan is defined as private but exported later.\nqiitan := func() {\n        return \"Hi! How are you?\"\n}\n\n// Export functions as public.\nexport {\n    // world functions returns a formatted string including the givn args as\n    // \"Hello, <arg> <arg>...!\".\n    world: func(...args) {\n\n        who := join(args...) // pass all the args through\n\n        return \"Hello, \" + who + \"!\"\n    },\n\n    // qiitan fuction returns a greeting message. Which was defined previously.\n    qiitan: qiitan\n}\n",
}
