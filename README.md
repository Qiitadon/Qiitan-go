<!-- markdownlint-disable MD041 -->
> - <sub><sup><a href="https://github.com/increments/mastodon/blob/qiitadon/app/javascript/images/qiitadon-getting-started.png">キーたん（Qiitan）</a>は、<a href="https://qiita.com/">Qiita</a><sup>™️</sup> の SNS である「<a href="https://qiitadon.com/">Qiitadon</a><sub><sup><sup>β</sup></sup></sub>」のマスコット・キャラクターです。</sup></sub>
> - <sub><sup>`Qiitan-go` は非公認の Qiitan のファン・アプリです。<a href="https://qiita.com/">Qiita</a><sup>™️</sup> とは一切関係がありません。</sup></sub>

---

[![platform icon](https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-blue "win,mac,linux")](https://github.com/Qithub-BOT/Qiitan-go/releases/latest "view latest releases page")
[![homebrew icon](https://img.shields.io/badge/homebrew-macos%20%7C%20linux-blue "win,mac,linux")](https://github.com/Qithub-BOT/Qiitan-go#install "view latest releases page")
[![MIT license](https://img.shields.io/github/license/Qithub-BOT/Qiitan-go)](https://github.com/Qithub-BOT/Qiitan-go/blob/main/LICENSE "view license page")

- [`qiitan` インタプリタのインストール or ダウンロード](https://github.com/Qithub-BOT/Qiitan-go#install)

# Qiitan-go<sub><sup><sup>alpha</sup></sup></sub>

[キーたん語（Qiitan-go）](https://github.com/Qiitadon/Qiitan-go)は、[キーたん（Qiitan）](https://github.com/increments/mastodon/blob/qiitadon/app/javascript/images/qiitadon-getting-started.png) ファンの、Qiitan ファンによる、Qiitan ファンのための「お遊び用スクリプト言語」です。

## `qiitan` スクリプトの Hello World

### スクリプトと実行例

```shellsession
$ # スクリプト
$ cat ./hello_world.qiitan
/* Hello World of Qiitan Script */
foo := import("fmt")
foo.println("Hello World!")

$ # 実行
$ qiitan ./helloworld.qiitan
Hello World!
```

- スクリプトの簡易説明（`./hello_world.qiitan`）
    ```go
    foo := import("fmt")
    foo.println("Hello World!")
    ```
    - この qiitan スクリプトは、標準モジュール（`qiitan` インタプリタに同梱されているモジュール）から `fmt` モジュールを `foo` に代入し、その `fmt` モジュール内で定義されている `println` 関数（改行付き `print` 関数）を呼び出して標準出力に出力しています。

### 対話モード（[REPL](https://ja.wikipedia.org/wiki/REPL)）での実行例

```shellsession
$ # 対話モード（REPL）で実行
$ qiitan
Welcome to Qiitan Script Interactive Mode! (To exit just type `bye`)
> foo := import("fmt")
LOG: {__module_name__: "fmt", print: <user-function>, printf: <user-function>, println: <user-function>, sprintf: <user-function>}
> foo.println("Hello World!")
Hello World!
LOG: <undefined> object returned. Perhaps the statement has no return or nothing is assigned to it.
>
> bye
Nice chatting with you. Thank you! Bye-bye~.
$
```

<details><summary>上記の簡易説明</summary><br>

対話モードでは、出力した内容および代入した内容が適宜表示されます。

例えば、`foo := import("fmt")` で `foo` 変数に `fmt` モジュールを代入した場合、続く `LOG:` の内容から、`foo` 変数には `fmt` モジュールの中身である `print()`, `printf()`, `println()`, `sprintf()` の関数が定義されていることが確認できます。

```shellsession
> foo := import("fmt")
LOG: {__module_name__: "fmt", print: <user-function>, printf: <user-function>, println: <user-function>, sprintf: <user-function>}
```

次に `foo.println("Hello World!")` で標準出力に "`Hello World!`" を出力していますが、続く `LOG: <undefined> object returned` は、`fmt.println()` に戻り値がなかった（型がない値が返ってきた）ことを表しています。

```shellsession
> foo.println("Hello World!")
Hello World!
LOG: <undefined> object returned. Perhaps the statement has no return or nothing is assigned to it.
```

これは、`qiitan` スクリプトでは `return` のない関数は `<undefined>` オブジェクト（型がないことを示すオブジェクト）が自動的に返されるためです。

---

</details>

### プリ・コンパイルと実行

キーたん語は、なんちゃってプリ・コンパイル型の言語です。

`qiitan` スクリプトの実行だけでなく、その中間ファイル（バイトコード）を出力することができ、この中間ファイルも `qiitan` インタプリタで実行することができます。
`qiitan` スクリプトが複数ファイルに分割されている場合、1 つのファイルとして配布したい場合に便利です。（速度は大して変わりません）

```shellsession
$ qiitan -compile ./helloworld.qiitan
compiled: ./helloworld.qtn

$ qiitan ./helloworld.qtn
Hello World!
```

---

## リソース

- [`qiitan` スクリプトのサンプル集](./samples)
- [`qiitan` スクリプトのドキュメント](https://Qiitadon.github.io/Qiitan-go/)

---

## Install

- [Homebrew](https://brew.sh/index_ja) (**macOS:** x86_64/Intel/AMD64, M1/ARM64, **Linux & Windows WSL2:** x86_64, ARM64, ARM v6)

    ```bash
    brew install Qiitadon/apps/qiitan
    ```

- 手動インストール（Windows, macOS, Linux）
    - [Releases ページ](https://github.com/Qiitadon/Qiitan-go/releases/latest)
    - `qiitan` インタプリタは単体バイナリです。リリース・ページから OS + アーキテクチャに該当するバイナリをダウンロードし、実行権限を与えてパスの通ったディレクトリに設置するだけで利用できます。
    - macOS で「`Operation not permitted`」エラーが出る場合は、ターミナル・アプリにディスク・アクセスの権限が与えられていない可能性があります。
        - [macOS の「Operation not permitted」を回避する（du/ls/mv/cp 実行時）](https://qiita.com/KEINOS/items/0366f1c281b574a79cfb) @ Qiita

---

## Statuses

このリポジトリでは、以下のような最低限のセキュリティ対策が施されています。

[![go1.14+](https://github.com/Qiitadon/Qiitan-go/actions/workflows/go-versions.yml/badge.svg)](https://github.com/Qiitadon/Qiitan-go/actions/workflows/go-versions.yml "Unit Tests on Go 1.16, 17 and latest")
[![Platform Tests](https://github.com/Qiitadon/Qiitan-go/actions/workflows/platform-test.yml/badge.svg)](https://github.com/Qiitadon/Qiitan-go/actions/workflows/platform-test.yml "Test on Win, macOS, Linux")
[![golangci-lint](https://github.com/Qiitadon/Qiitan-go/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Qiitadon/Qiitan-go/actions/workflows/golangci-lint.yml "Static Analysis")
[![CodeQL](https://github.com/Qiitadon/Qiitan-go/actions/workflows/codeQL-analysis.yml/badge.svg)](https://github.com/Qiitadon/Qiitan-go/actions/workflows/codeQL-analysis.yml "Vulnerability Scan")
[![codecov](https://codecov.io/gh/Qiitadon/Qiitan-go/branch/main/graph/badge.svg?token=uW30s2bK8M)](https://codecov.io/gh/Qiitadon/Qiitan-go "Code Coverage")
[![Go Report Card](https://goreportcard.com/badge/github.com/Qiitadon/Qiitan-go)](https://goreportcard.com/report/github.com/Qiitadon/Qiitan-go "Code Quality")

---

## 言語拡張開発（コントリビュート）

[![go1.16+](https://img.shields.io/badge/Go-1.16+-blue?logo=go)](https://github.com/Qiitadon/Qiitan-go/actions/workflows/go-versions.yml "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/Qiitadon/Qiitan-go.svg)](https://pkg.go.dev/github.com/Qiitadon/Qiitan-go)

- `qiitan` コマンド（インタプリタ）自体は Go 言語で開発されています。
- `qiitan` コマンド（インタプリタ）は、スタンドアロン（単体）のバイナリとして動作するように設計されています。
- [CONTRIBUTING](./.github/CONTRIBUTING.md)

---

## License

- [MIT](https://github.com/Qiitadon/Qiitan-go/LICENSE) License.
- Copyright: © 2022 [The Qiitan-go Contributors](https://github.com/Qiitadon/Qiitan-go/graphs/contributors).

### 謝辞

- `qiitan` スクリプトは [Tengo 言語](https://github.com/d5/tengo)をベースに拡張した上位互換言語です。
- `qiitan` インタプリタは、以下の Go モジュールにお世話になっています。
  - [go.mod](./go.mod)
