<!-- Code generated using /gen/gen_readme.sh; DO NOT EDIT. -->
# Qiitan スクリプトの言語仕様

- [Qiitan-go リポジトリ](https://github.com/Qiitadon/Qiitan-go) @ GitHub

## 目次

* [スクリプト・ファイル](./specs/001_language_syntax/001_scrip_file.md#スクリプトファイル)
  * [ファイル・フォーマット](./specs/001_language_syntax/001_scrip_file.md#ファイルフォーマット)
    * [文字コード](./specs/001_language_syntax/001_scrip_file.md#文字コード)
    * [改行コード](./specs/001_language_syntax/001_scrip_file.md#改行コード)
    * [拡張子](./specs/001_language_syntax/001_scrip_file.md#拡張子)
  * [ファイルのアクセス権](./specs/001_language_syntax/001_scrip_file.md#ファイルのアクセス権)
  * [ファイルの実行](./specs/001_language_syntax/001_scrip_file.md#ファイルの実行)

* [型（type）](./specs/001_language_syntax/002_type.md#型type)
  * [Qiitan スクリプトで使える型の一覧](./specs/001_language_syntax/002_type.md#qiitan-スクリプトで使える型の一覧)

* [文 (statement)](./specs/001_language_syntax/003_statement.md#文-statement)

* [式 (expression)](./specs/001_language_syntax/004_expression.md#式-expression)

* [値 (value)](./specs/001_language_syntax/005_value.md#値-value)

* [immutable](./specs/001_language_syntax/999_values_and_types.md#immutable)
* [error](./specs/001_language_syntax/999_values_and_types.md#error)
* [undefined](./specs/001_language_syntax/999_values_and_types.md#undefined)

* [組み込み関数（Built\-in Functions）](./specs/002_builtin_functions.md#組み込み関数built-in-functions)

* [error type（エラー型と値）](./specs/002_functions/011_type_error.md#error-typeエラー型と値)
  * [エラーの作成](./specs/002_functions/011_type_error.md#エラーの作成)
  * [エラー内容の取り出し](./specs/002_functions/011_type_error.md#エラー内容の取り出し)
  * [エラー型の確認](./specs/002_functions/011_type_error.md#エラー型の確認)

## モジュール

### Tengo 互換モジュール

* [base64](./modules_tengo/base64.md#base64)
  * [encode](./modules_tengo/base64.md#encode)
  * [decode](./modules_tengo/base64.md#decode)
  * [url\_encode](./modules_tengo/base64.md#url_encode)
  * [url\_deocde](./modules_tengo/base64.md#url_deocde)
  * [raw\_encode](./modules_tengo/base64.md#raw_encode)
  * [raw\_decode](./modules_tengo/base64.md#raw_decode)
  * [raw\_url\_encode](./modules_tengo/base64.md#raw_url_encode)
  * [raw\_url\_decode](./modules_tengo/base64.md#raw_url_decode)
* [enum](./modules_tengo/enum.md#enum)
  * [all](./modules_tengo/enum.md#all)
  * [any](./modules_tengo/enum.md#any)
  * [chunk](./modules_tengo/enum.md#chunk)
  * [at](./modules_tengo/enum.md#at)
  * [each](./modules_tengo/enum.md#each)
  * [filter](./modules_tengo/enum.md#filter)
* [fmt](./modules_tengo/fmt.md#fmt)
  * [print](./modules_tengo/fmt.md#print)
  * [println](./modules_tengo/fmt.md#println)
  * [printf](./modules_tengo/fmt.md#printf)
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

## 用語集

「キーたん語」の言語仕様もしくは「qiitan スクリプト」の構文など、仕様の中で使われる用語の定義と説明です。（順不順）

* [enumerable](./index/-able.md#enumerable)
* [iterable](./index/-able.md#iterable)
* [mutable/immutable](./index/-able.md#mutableimmutable)
* [組み込み関数](./index/built_in_functions.md#組み込み関数)
* [バイトコードとコンパイル](./index/bytecode_compile.md#バイトコードとコンパイル)
* [qiitan インタプリタ](./index/interpreter.md#qiitan-インタプリタ)
* [モジュール](./index/module.md#モジュール)
  * [標準モジュール](./index/module.md#標準モジュール)
  * [外部モジュール](./index/module.md#外部モジュール)
* [オブジェクト](./index/object.md#オブジェクト)
* [キーたん語（qiitan スクリプト）](./index/qiitanlang.md#キーたん語qiitan-スクリプト)
  * [「キーたん語」と「qiitan スクリプト」の違い](./index/qiitanlang.md#キーたん語とqiitan-スクリプトの違い)
  * [Qiitan と qiitan の使い分け](./index/qiitanlang.md#qiitan-と-qiitan-の使い分け)
* [tengo VM（tengo 仮想マシン）](./index/tengo_vm.md#tengo-vmtengo-仮想マシン)
* [Tengo 言語](./index/tengolang.md#tengo-言語)

---

- View the repository: [Qiitan-go](https://github.com/Qiitadon/Qiitan-go) @ GitHub
- Table of contents created by [gh-md-toc](https://github.com/ekalinin/github-markdown-toc.go) @ GitHub
