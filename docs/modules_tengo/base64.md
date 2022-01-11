# base64

`base64` は base64 方式のエンコーダー、デコーダーのモジュールです。

```go
base64 := import("base64")
```

```go
fmt := import("fmt")
base64 := import("base64")
str := "Hello, Qiitan!"

fmt.println(base64.encode(str))

// Output: SGVsbG8sIFFpaXRhbiE=
```

## decode

`eoncode` （MIME / RFC 4648 準拠の base64 パディングあり）でエンコードされた `string` をデコードします。

```go
decode(string) bytes
```

```go
base64 := import("base64")
raw := base64.decode("SGVsbG8sIFFpaXRhbiE=")
// raw: Hello, Qiitan!
```

## encode

指定した `bytes` を MIME / RFC 4648 準拠の base64 パディングありでエンコードします。

```go
encode(bytes) string
```

```go
base64 := import("base64")
enc := base64.encode("Hello, Qiitan!")
// enc: SGVsbG8sIFFpaXRhbiE=
```

## raw_decode

`raw_eoncode` （MIME / RFC 4648 準拠の base64 パディングなし）でエンコードされた `string` をデコードします。

```go
raw_decode(string) bytes
```

```go
base64 := import("base64")
raw := base64.raw_decode("SGVsbG8sIFFpaXRhbiE")
// raw: Hello, Qiitan!
```

## raw_encode

指定した `bytes` を MIME / RFC 4648 準拠の base64 パディングなしでエンコードします。

```go
raw_encode(bytes) string
```

```go
base64 := import("base64")
enc := base64.raw_encode("Hello, Qiitan!")
// enc: SGVsbG8sIFFpaXRhbiE
```

## raw_url_decode

```go
raw_url_decode()
```

- `decode(s)`: returns the bytes represented by the base64 string s.
- `encode(src)`: returns the base64 encoding of src.
- `raw_decode(s)`: returns the bytes represented by the base64 string s which omits the padding.
- `raw_encode(src)`: returns the base64 encoding of src but omits the padding.
- `raw_url_decode(s)`: returns the bytes represented by the url-base64 string s which omits the padding.
- `raw_url_encode(src)`: returns the url-base64 encoding of src but omits the padding.
- `url_decode(s)`: returns the bytes represented by the url-base64 string s.
- `url_encode(src)`: returns the url-base64 encoding of src.

> [目次に戻る ↩️](../README.md)

---
