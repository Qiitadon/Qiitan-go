<!-- Code generated using /gen/gen_readme.sh; DO NOT EDIT. -->
# Qiitan スクリプトの言語仕様

- [Qiitan-go リポジトリ](https://github.com/Qiitadon/Qiitan-go) @ GitHub

## 目次

* [Qiitan スクリプトの基本構文](./specs/001_language_syntax.md#qiitan-スクリプトの基本構文)
  * [スクリプトの文字コード](./specs/001_language_syntax.md#スクリプトの文字コード)
  * [型と値](./specs/001_language_syntax.md#型と値)
  * [文（statement）、式（expression）、値（value）](./specs/001_language_syntax.md#文statement式expression値value)

* [組み込み関数（Built\-in Functions）](./specs/002_builtin_functions.md#組み込み関数built-in-functions)

* [用語集](./specs/999_words_and_terms.md#用語集)
  * [キーたん語（qiitan スクリプト）](./specs/999_words_and_terms.md#キーたん語qiitan-スクリプト)
  * [qiitan インタプリタ](./specs/999_words_and_terms.md#qiitan-インタプリタ)
  * [Tengo 言語](./specs/999_words_and_terms.md#tengo-言語)
  * [モジュール](./specs/999_words_and_terms.md#モジュール)
  * [バイトコードとコンパイル](./specs/999_words_and_terms.md#バイトコードとコンパイル)
  * [tengo VM（tengo 仮想マシン）](./specs/999_words_and_terms.md#tengo-vmtengo-仮想マシン)
  * [オブジェクト](./specs/999_words_and_terms.md#オブジェクト)

* [error type（エラー型と値）](./specs/functions/011_type_error.md#error-typeエラー型と値)
  * [エラーの作成](./specs/functions/011_type_error.md#エラーの作成)
  * [エラー内容の取り出し](./specs/functions/011_type_error.md#エラー内容の取り出し)
  * [エラー型の確認](./specs/functions/011_type_error.md#エラー型の確認)

## モジュール

### Tengo 互換モジュール

* [base64](./modules_tengo/base64.md#base64)
  * [decode](./modules_tengo/base64.md#decode)
  * [encode](./modules_tengo/base64.md#encode)
  * [raw\_decode](./modules_tengo/base64.md#raw_decode)
  * [raw\_encode](./modules_tengo/base64.md#raw_encode)
  * [raw\_url\_decode](./modules_tengo/base64.md#raw_url_decode)
* [enum](./modules_tengo/enum.md#enum)
* [fmt](./modules_tengo/fmt.md#fmt)
* [hex](./modules_tengo/hex.md#hex)
* [json](./modules_tengo/json.md#json)
* [math](./modules_tengo/math.md#math)
* [os](./modules_tengo/os.md#os)
* [rand](./modules_tengo/rand.md#rand)
* [text](./modules_tengo/text.md#text)
* [times](./modules_tengo/times.md#times)

### qiitan 独自モジュール

* [hash](./modules_qiitan/hash.md#hash)
  * [md5](./modules_qiitan/hash.md#md5)
  * [sha1](./modules_qiitan/hash.md#sha1)
  * [sha224](./modules_qiitan/hash.md#sha224)
  * [sha256](./modules_qiitan/hash.md#sha256)
  * [sha384](./modules_qiitan/hash.md#sha384)
  * [sha512](./modules_qiitan/hash.md#sha512)

---

- View the repository: [Qiitan-go](https://github.com/Qiitadon/Qiitan-go) @ GitHub
- Table of contents created by [gh-md-toc](https://github.com/ekalinin/github-markdown-toc.go) @ GitHub
