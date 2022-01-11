# hash

`hash` は文字列用の[ハッシュ関数](https://ja.wikipedia.org/wiki/%E3%83%8F%E3%83%83%E3%82%B7%E3%83%A5%E9%96%A2%E6%95%B0)をまとめたモジュールです。

- インポート
    ```go
    hash := import("hash")
    ```
- 使用例
    ```go
    fmt := import("fmt")
    hash := import("hash")

    str := "Hello, Qiitan!"

    fmt.println(hash.md5(str))
    // Output: 82ed57a748cc18329b954105429f49f5
    ```

- ハッシュ関数や、ハッシュに使われるアルゴリズムについては、以下の Qiita 記事を参考にしてください。
  - [hashアルゴリズムとハッシュ値の長さ一覧（＋ハッシュ関数の基本と応用）](https://qiita.com/KEINOS/items/c92268386d265042ea16) @ Qiita

## md5

`md5` 関数は、引数の文字列を [MD5](https://ja.wikipedia.org/wiki/MD5) アルゴリズムを使ったハッシュ値を返します。

```go
md5(string) string
```

```go
hash := import("hash")
result := hash.md5("Hello, Qiitan!=")
// result: 82ed57a748cc18329b954105429f49f5
```

> [目次に戻る ↩️](../README.md)

---

## sha1

`sha1` 関数は、引数の文字列を [SHA1](https://ja.wikipedia.org/wiki/SHA-1) アルゴリズムを使ったハッシュ値を返します。

```go
sha1(string) string
```

```go
hash := import("hash")
result := hash.sha1("Hello, Qiitan!")
// result: 2343d74f9436884f362d8a753fa11c153d80f0c9
```

> [目次に戻る ↩️](../README.md)

---

## sha224

`sha224` 関数は、引数の文字列を [SHA2](https://ja.wikipedia.org/wiki/SHA-2) の SHA-224 アルゴリズムを使ったハッシュ値を返します。

```go
sha224(string) string
```

```go
hash := import("hash")
result := hash.sha224("Hello, Qiitan!")
// result: cba7426676e50fe569b21c28aa2294893bfa03967ffb3601ab605f6c
```

> [目次に戻る ↩️](../README.md)

---

## sha256

`sha256` 関数は、引数の文字列を [SHA2](https://ja.wikipedia.org/wiki/SHA-2) の SHA-256 アルゴリズムを使ったハッシュ値を返します。

```go
sha256(string) string
```

```go
hash := import("hash")
result := hash.sha256("Hello, Qiitan!")
// result: fa7fcdc946b8a96ba78424293a0af52b2d60cfecd8de0e03857edd4a005147f1
```

> [目次に戻る ↩️](../README.md)

---

## sha384

`sha256` 関数は、引数の文字列を [SHA2](https://ja.wikipedia.org/wiki/SHA-2) の SHA-384 アルゴリズムを使ったハッシュ値を返します。

```go
sha384(string) string
```

```go
hash := import("hash")
result := hash.sha384("Hello, Qiitan!")
// result: aeab0e947bcdfe2ca2e1262453dd232dd55bc3c700e00a02a79cb032045be3ac9e5cd4e989416f14763d0283a773c5fd
```

> [目次に戻る ↩️](../README.md)

---

## sha512

`sha256` 関数は、引数の文字列を [SHA2](https://ja.wikipedia.org/wiki/SHA-2) の SHA-512 アルゴリズムを使ったハッシュ値を返します。

```go
sha512(string) string
```

```go
hash := import("hash")
result := hash.sha512("Hello, Qiitan!")
// result: 123163af3c5575310d56ca1cd7d8b6253f072f5ed554ec63659e8f7627359742e1b621cd3a10fb95e72ce03d2deacbb94efbc11db8768d394b6eadd1606203a2
```

> [目次に戻る ↩️](../README.md)

---
