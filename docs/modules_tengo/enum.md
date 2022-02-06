![](https://img.shields.io/badge/Tengo%20Compatible-Yes-blue "Tengo compatible module")
![](https://img.shields.io/badge/Module%20type-Script-blue "Module made in Tengo Script")

# enum

`enum` モジュールは、可算（数え上げることのできる値、`enumerable`）に関する関数群です。主に配列の繰り返し処理に関するものです。

```go
enum := import("enum")
```
```go
enum := import("enum")

validator := func(key, value){
    return value < 6
}

arr := [1, 2, 3, 4, 5]

result := enum.all(arr, validator)
// result --> true
```

- [このモジュールの定義ソース](https://github.com/d5/tengo/blob/master/stdlib/srcmod_enum.tengo) @ github.com/d5/tengo

---

## all

```go
all(x, fn) bool
// x: array type, map type, immutable array type, immutable map type
// fn: function type
```

`all` は、配列のバリデーション（検証）用関数です。
配列 `x` のすべての要素を `function` 型 `fn` の関数で検証し、すべての結果が `true` だった場合に `true` を返します。

- 第 1 引数の `x` が配列（反復処理が可能な型）ではない場合、`undefined` 型が返されます。
- 第 2 引数の `fn` は、`function` 型である必要があります。
    - 検証される要素のインデックスもしくはキーと、その値が渡されるので、検証結果の真偽（`bool`）を返す必要があります。
    ```go
    function(key, value) bool
    // key: int type, string type
    // value: object type
    ```

```go
enum := import("enum")

// 要素の値が 6 より小さいかのバリデーション
validator := func(key, value){
    return value < 6
}

arr := [1, 2, 3, 4, 5]

// すべての値が 6 より小さいかバリデート
result := enum.all(arr, validator)
// result --> true
```

> [目次に戻る ↩️](../)

---

## any

```go
any(x, fn) bool
// x: array type, map type, immutable array type, immutable map type
// fn: function type
```

`any` は、配列のバリデーション（検証）用関数です。
配列 `x` のすべての要素を `function` 型 `fn` の関数で検証し、1 つでも結果が `true` だった場合に `true` を返します。

- 第 1 引数の `x` が配列（反復処理が可能な型）ではない場合、`undefined` 型が返されます。
- 第 2 引数の `fn` は、`function` 型である必要があります。
    - 検証される要素のインデックスもしくはキーと、その値が渡されるので、検証結果の真偽（`bool`）を返す必要があります。
    ```go
    function(key, value) bool
    ```

```go
enum := import("enum")

// 要素の値が 3 であるかのバリデーション
validator := func(key, value){
    return value == 3
}

arr := [1, 2, 3, 4, 5]

// 配列の値に 3 が含まれているかバリデート
result := enum.any(arr, validator)
// result --> true
```

> [目次に戻る ↩️](../)

---

## chunk

```go
chunk(x, size) [object]
// x: array type, map type, immutable array type, immutable map type
// size: int type
```

`chunk` は、配列の分割関数です。
配列 `x` を、`size` の長さのチャンク（グループ、塊）に分割し多次元配列を返します。
`x` が均等に分割できない場合は、最後のチャンクが残りの要素になります。

- 第 1 引数の `x` が配列（反復処理が可能な型）ではない場合、`undefined` 型が返されます。

```go
enum := import("enum")

arr := [1, 2, 3, 4, 5, 6, 7]
size := 5

result := enum.chunk(arr, size)
// result --> [[1, 2, 3, 4, 5], [6, 7]]
```

> [目次に戻る ↩️](../)

---

## at

```go
at(x, key) object
// x: array type, map type, immutable array type, immutable map type
// key: int type
```

`at` は、配列の抽出関数です。
`x` が連想配列（`array`）の場合、指定されたインデックスにある要素を返します。
`x` が添字配列（`map`）の場合、指定されたキーにある要素を返します。

- 第 1 引数の `x` が配列型（反復処理が数え上げができる型）ではない場合、`undefined` 型が返されます。
- この関数は元の配列を操作しません（返された要素は元の配列から削除されません）。

```go
enum := import("enum")

arr := ["zero", "one", "two", "three", "four", "five"]
index := 3

result := enum.at(arr, index)
remain := arr
// result --> three
// remain --> ["zero", "one", "two", "three", "four", "five"]
```

> [目次に戻る ↩️](../)

---

## each

```go
each(x, fn)
// x: array type, map type, immutable array type, immutable map type
// fn: function type
```

`each` は、配列の反復処理（イテレーション）関数です。
配列 `x` のすべての要素に対して `function` 型 `fn` の関数を実行します。

- 第 1 引数の `x` が配列（反復処理が可能な型、`iterable`）ではない場合、`undefined` 型が返されます。
- 第 2 引数の `fn` は、`function` 型である必要があります。
    - 検証される要素のインデックスもしくはキーと、値が渡されるので、検証結果の真偽（`bool`）を返す必要があります。
    ```go
    function(key, value) bool
    // key: int type, string type
    // value: object type
    ```
- **【重要】**
  - **`x` が、添字配列（`map` 型）か連想配列（`array` 型）かで動作が異なる**ので気をつけてください。

    **`x` が「`map` 型」の場合、イテレーションの順番はランダムになります**。つまり、要素の並び順が基本的に変動します。逆に `x` が連想配列（`array`）の場合は、元の要素の並び順でイテレーションされます。

    これは、Go 言語の設計思想（[仕様](https://go.dev/ref/spec#For_range)）によるものです。`map` の場合は順不同である（キー名でアクセスする）ことが前提であること以外に、Go 言語の「`map` の要素数が多い（8 要素以上の） 場合、高速化のため[低レベルの命令](https://ja.wikipedia.org/wiki/%E4%BD%8E%E6%B0%B4%E6%BA%96%E8%A8%80%E8%AA%9E)を利用する」仕様に起因します。

    低レベルの命令を使うということは、バイナリをビルドした時のコンパイラの違いによって挙動が異なるということでもあります。そして、命令の種類によっては（`map` の場合に）高速処理のために要素の並び順が変わるものがあります。そのため、[その仕様に気づかずに「並び順に依存した `map` のコード」が、予期せぬエラーや不要な問い合わせの温床となっていた](https://github.com/golang/go/issues/6719)ことが問題になりました。

    そこで、テストで 8 要素以下の `map` のイテレーションを書いても並び順に依存していることに気付けるように、**Go 言語では「`map` を `for 〜 range`（Go 言語の反復命令）を使ってイテレーションを行う場合は、あえてランダムにする」という仕様になっています**。

```go
enum := import("enum")

arr := [1, 2, 3, 4, 5, 6, 7]

// すべての値を 2 倍にする
walk := func(k, v) {
    arr[k] = v*2
}

enum.each(arr, walk)

// arr --> [2, 4, 6, 8, 10, 12, 14]
```
```go
enum := import("enum")

// すべての値を 2 倍にする
walk = func(k, v) {
    map[k] = v*2
}

map := {"one": 1, "two": 2, "three": 3, "four": 4}

enum.each(map, walk)

// 並び順が変わっていることに注目
// map --> {two: 4, three: 6, four: 8, one: 2}
```

> [目次に戻る ↩️](../)

---

## filter

```go
filter(x, fn) [object]
// x: array type, map type, immutable array type, immutable map type
// fn: function type
```

`filter` は、配列の抽出関数です。
配列 `x` のすべての要素を `fn` の関数（`function` 型）で検証し `true` だった場合の要素を配列で返します。

- 第 1 引数の `x` が配列（繰り返し可能な型、`iterable`）ではない場合、`undefined` 型が返されます。
- 第 2 引数の `fn` は、`function` 型である必要があります。
  - 検証される要素のインデックスキーと、その値が渡されるので、検証結果の真偽（`bool`）を返す必要があります。
    ```go
    function(key, value) bool
    // key: int type, string type
    // value: object type
    ```
  - **【重要】**
    `each` 同様、**`x` が「`map` 型」の場合、イテレーションの順番はランダムになります**。つまり、要素の並び順が基本的に変動します。逆に `x` が連想配列（`array`）の場合は、元の要素の並び順でイテレーションされます。この仕様の詳細は [`each`](#each) をご覧ください。

```go
enum := import("enum")

// 偶数フィルター
filter := func(k, v) {
    return v % 2 == 0
}

arr := [1, 2, 3, 4, 5, 6, 7]

result := enum.filter(arr, filter)

// result --> [2, 4, 6]
```

> [目次に戻る ↩️](../)

---
