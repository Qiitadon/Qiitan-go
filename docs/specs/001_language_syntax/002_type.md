# 型（type）

キーたん語はオブジェクト指向型ではありませんが、全てがオブジェクトであり、「型を持った値」です。

```golang
19 + 84               // int 型
"aomame" + `kawa`     // string 型
-9.22 + 1e10          // float 型
true || false         // bool 型
'九' > '9'            // char 型（Go でいう rune と同等）
[1, false, "foo"]     // array 型（連想配列）
{a: 12.34, b: "bar"}  // map 型（添字配列）
func() { /*...*/ }    // function 型
```

関数も「`function` という型を持った値」として扱います。例えば、ユーザ関数を作成したい場合は、変数に代入して使います。いわゆる[無名関数](https://ja.wikipedia.org/wiki/%E7%84%A1%E5%90%8D%E9%96%A2%E6%95%B0)と同じ使い方です。

```golang
hoge := func() { /* ... ここに処理 ... */ } // hoge は function 型の値

fuga := hoge() // fuga には hoge の実行結果（func() 内の return の値）が代入される
piyo := hoge   // piyo には hoge の関数の中身がコピーされ、piyo もユーザ関数となる
```

同様に、スクリプト内で使いたいモジュール（関数群）も変数に代入して利用します。
以下は、標準出力系の `fmt` モジュールの利用例です。モジュール名と変数名は同名にするのが慣習ですが、別名でも構いません。

```golang
fmt := import("fmt") // モジュールの割り当て

fmt.println("foo bar") // モジュール内の関数の利用
```

## Qiitan スクリプトで使える型の一覧

| 値の型<br>（タイプ名） | 値の説明 | Go の場合の型 |
| :---: | :---- | :---: |
| int | [符号付き](https://ja.wikipedia.org/wiki/%E7%AC%A6%E5%8F%B7%E4%BB%98%E6%95%B0%E5%80%A4%E8%A1%A8%E7%8F%BE) 64 ビットの整数の値 | `int64` |
| float | 64 ビットの浮動小数 | `float64` |
| bool | `true`（真） `false`（偽） の[真理値](https://ja.wikipedia.org/wiki/%E7%9C%9F%E7%90%86%E5%80%A4) | `bool` |
| char | Unicode の 1 文字ぶんのバイト配列<br>（1 コードポイント） | `rune` |
| string | Unicode 文字列 | `string` |
| bytes | 1 バイト（8 ビット）ごとの配列 | `[]byte` |
| error | エラーを示す値 | - |
| time | 時間を示す値 | `time.Time` |
| array | 連想配列 | `[]interface{}` |
| immutable array | 定数化された連想配列<br>（値の変更不可） | - |
| map | 添字付き配列 | `map[string]interface{}` |
| immutable map | 定数化された添字配列<br>（値の変更不可） | - |
| undefined | 型が未定義の値 | - |
| function | 関数を示す値 | - |
| _user-defined_ | ユーザ定義型の値 | - |

> [目次に戻る ↩️](../)
