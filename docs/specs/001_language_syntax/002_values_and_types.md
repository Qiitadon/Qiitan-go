# 値と型

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

---

### immutable

値には、不変なもの（イミュータブル）と可変なもの（ミュータブル）があります。

不変な値とは、生成した時点の状態を後から変更できないものを言います。いわゆる「定数」のようなもので、定義後に値を変更できないようなものです。

キーたん語では、**連想配列（`array`）と添字配列（`map`）を除き、基本的に全てが不変（イミュータブル）な値です**。

```go
s := "12345"   // s のオブジェクトに string 型の値を代入する
s[1] = 'b'     // NG: 文字列型は不変です

a := [1, 2, 3] // a --> array 型
a[1] = "two"   // OK: `a` は `[1, "two", 3]` になります
```

`array` 型や `map` 型の値を不変にしたい場合は、`immutable()` を使って不変化させたものを使います。

```go
b := immutable([1, 2, 3])
b[1] = "foo"   // NG: 'b' は不変化された array が入ったオブジェクトです
```

とは言え、キーたん語は、そこまで「型々」言う言語ではありません。新しい値を既存の変数（オブジェクト）に代入することは可能です。

```go
s := "abc"
s = "foo"                  // OK
a := immutable([1, 2, 3])
a = false                  // OK
```

なお、不変の値を `copy()` を使ってコピーすると、「可変」のコピーが返されます。

```go
a := immutable([1, 2, 3])
a[1] = "foo"   // NG: 'a' は不変化された array が入ったオブジェクトです

b := copy(a)
b[1] = "foo"   // OK: `b` は `[1, "foo", 3]` になります
```

また，連想配列（`array`）や添字配列（`map`）の値の個々の要素には，明示的に不変性が設定されていない限り，不変性は適用されません。

```go
a := immutable({b: 4, c: [1, 2, 3]})
a.b = 5        // NG: 'a.b' は string で不変なため
a.c[1] = 5     // OK: 'a.c' は不変化されていないため

a = immutable({b: 4, c: immutable([1, 2, 3])})
a.c[1] = 5     // NG
```

> [目次に戻る ↩️](../)

---

### error

キーたん語では、エラーは "`error`" という型の値で表されます。

エラーは `error()` を使用して作成され、基礎となる型の値を持たなければなりません。エラー値の基礎となる値は、`.value` セレクタを使ってアクセスできます。

```go
err1 := error("oops")    // string 型の値を持つエラー型
err2 := error(1+2+3)     // int 型の値を持つエラー型

// 'is_error' は組み込み関数で、引数がエラー型の場合に true を返します。
if is_error(err1) {
  err_val := err1.value  // err1 に内在する値を取得し err_val に代入
}
```

> [目次に戻る ↩️](../)

---

### undefined

キーたん語では、"`undefined`" 型の値は、予想外の値や存在しない値を表すのに使われます。

- 値を返さない関数は、明示的に `undefined` の値を返すとみなされます。
- `array` や `map` などの複合値型で、キーやインデックスが存在しない場合、`undefined` を返すことがあります。
- 型変換の組み込み関数で、デフォルト値がない場合は、変換に失敗すると `undefined` を返します。

```go
a := func() { b := 4 }()    // a == undefined
b := [1, 2, 3][10]          // b == undefined
c := {a: "foo"}["b"]        // c == undefined
d := int("foo")             // d == undefined
```

> [目次に戻る ↩️](../)

---

---

> [目次に戻る ↩️](../)
