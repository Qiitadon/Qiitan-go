![](https://img.shields.io/badge/Tengo%20Compatible-Yes-blue "Tengo compatible module")
![](https://img.shields.io/badge/Module%20type-Go-blue "Module made in Go")

# base64

`base64` は [RFC 4648](https://www.rfc-editor.org/rfc/rfc4648.html) に準拠した Base64 エンコーダー、デコーダーのモジュールです。

画像や音声ファイルといったバイナリ・データを、テキスト・データとして扱うことができます。

- エンコーダーは、バイナリデータ（`bytes`）を 64 進数の文字列に変換します。
  - 引数が `string` 型の場合は `bytes` 型に変換してバイナリデータとして扱います。
- デコーダーは、エンコードされたデータを `bytes` データに逆変換します。
- エンコードの末尾に付くパディング（`=`）を付けたくない場合は `raw_*` 系の関数を使います。

```go
base64 := import("base64")
```

```go
fmt := import("fmt")
base64 := import("base64")

// エンコード
fmt.println(base64.encode("Hello, Qiitan!"))
// デコード
fmt.println(base64.decode("SGVsbG8sIFFpaXRhbiE="))

// Output:
// SGVsbG8sIFFpaXRhbiE=
// Hello, Qiitan!
```

- [このモジュールの定義ソース](https://github.com/d5/tengo/blob/master/stdlib/base64.go) @ github.com/d5/tengo

---

## encode

```go
encode(x) string
// x: string type, bytes type
```

`encode` は、`x` を [RFC 4648](https://www.rfc-editor.org/rfc/rfc4648.html) 準拠の Base64 パディングありでエンコードします。

```go
base64 := import("base64")
enc := base64.encode("012abcABC _/*=~.+-\\")
// enc: MDEyYWJjQUJDIF8vKj1+ListXA==
```

## decode

```go
decode(x) bytes
// x: string type
```

`decode` は、`encode`（[RFC 4648](https://www.rfc-editor.org/rfc/rfc4648.html) 準拠の Base64 パディングあり）でエンコードされた `x` をデコードします。

```go
base64 := import("base64")
raw := base64.decode("MDEyYWJjQUJDIF8vKj1+ListXA==")
// raw: 012abcABC _/*=~.+-\
```

## url_encode

```go
url_encode(x) string
// x: string type, bytes type
```

`url_encode` は、`x` を [RFC 4648](https://www.rfc-editor.org/rfc/rfc4648.html) 準拠の URL Base64 パディングありでエンコードします。

`url_encode` は `encode` と似ていますが、エンコード文字に `+` と `/` の代わりに `-` と `_` を使い URL やファイル名でも使えるように配慮されています。

```go
base64 := import("base64")
enc := base64.url_encode("012abcABC _/*=~.+-\\")
// enc: MDEyYWJjQUJDIF8vKj1-ListXA==
```

## url_deocde

```go
url_decode(x) bytes
// x: string type
```

`url_deocde` は、`url_encode`（[RFC 4648](https://www.rfc-editor.org/rfc/rfc4648.html) 準拠の Base64 パディングあり）でエンコードされた `x` をデコードします。

```go
base64 := import("base64")
raw := base64.url_deocde("MDEyYWJjQUJDIF8vKj1-ListXA==")
// raw: 012abcABC _/*=~.+-\
```

## raw_encode

```go
raw_encode(x) string
// x: string type, bytes type
```

`raw_encode` は、`x` を [RFC 4648](https://www.rfc-editor.org/rfc/rfc4648.html) 準拠の Raw Base64 パディングなしでエンコードします。

`raw_encode` は `encode` と似ていますがエンコード結果にパディング（`=`）が付きません。

```go
base64 := import("base64")
enc := base64.raw_encode("012abcABC _/*=~.+-\\")
// enc: MDEyYWJjQUJDIF8vKj1+ListXA
```

## raw_decode

```go
raw_decode(x) bytes
// x: string type
```

`raw_decode` は、`raw_eoncode` （[RFC 4648](https://www.rfc-editor.org/rfc/rfc4648.html) 準拠の Raw Base64 パディングなし）でエンコードされた `x` をデコードします。

```go
base64 := import("base64")
raw := base64.raw_decode("MDEyYWJjQUJDIF8vKj1+ListXA")
// raw: 012abcABC _/*=~.+-\
```

## raw_url_encode

```go
raw_url_encode(x) string
// x: string type, bytes type
```

`raw_url_encode` は、`x` を [RFC 4648 準拠の URL Base64](https://www.rfc-editor.org/rfc/rfc4648.html#section-5) パディングなしでエンコードします。

`raw_url_encode` は `raw_encode` と似ていますが、エンコード文字に `+` と `/` の代わりに `-` と `_` を使い URL やファイル名でも使えるように配慮されています。

```go
base64 := import("base64")
enc := base64.raw_url_encode("012abcABC _/*=~.+-\\")
// enc: MDEyYWJjQUJDIF8vKj1-ListXA
```

## raw_url_decode

```go
raw_url_decode(x) bytes
// x: string type
```

`raw_url_decode` は、`raw_url_encode`（[RFC 4648 準拠の URL Base64](https://www.rfc-editor.org/rfc/rfc4648.html#section-5) パディングなし）でエンコードした `x` をデコードします。

```go
base64 := import("base64")
dec := base64.raw_url_decode("MDEyYWJjQUJDIF8vKj1-ListXA")
// dec: 012abcABC _/*=~.+-\
```

> [目次に戻る ↩️](../)

---
