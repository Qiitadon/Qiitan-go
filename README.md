<!-- markdownlint-disable MD041 -->
# Qiitan-go<sub><sup><sup>alpha</sup></sup></sub>

[キーたん語（Qiitan-go）](https://github.com/Qithub-BOT/Qiitan-go)は、Qiitan ファンの、Qiitan ファンによる、Qiitan ファンのための「お遊び用スクリプト言語」です。

> Qiitan は [Qiita](https://qiita.com/)™️ の SNS「[Qiitadon](https://qiitadon.com/)<sub><sup><sup>β</sup></sup></sub>」のマスコット・キャラクターです。
>
> - [キーたん（Qiitan）](https://github.com/increments/mastodon/blob/qiitadon/app/javascript/images/qiitadon-getting-started.png) @ Qiitadon
> - `Qiitan-go` は Qiitan のファン・アプリであり、[Qiita](https://qiita.com/)™️ とは一切関係がありません。

---

## `qiitan` スクリプトの Hello World

### スクリプトと実行例

- スクリプト例（`./helloworld.qiitan`）
    ```go
    foo := import("fmt")
    foo.println("Hello World!")
    ```
    - この qiitan スクリプトは、標準モジュール（`qiitan` インタプリタに同梱されているライブラリ）から `fmt` モジュールを `foo` に代入し、その `fmt` モジュール内で定義されている `println` 関数（改行付き `print` 関数）を呼び出して標準出力に出力しています。
- 実行例
    ```shellsession
    $ qiitan ./helloworld.qiitan
    Hello World!
    ```

### 対話モード（[REPL](https://ja.wikipedia.org/wiki/REPL)）での実行例

```shellsession
$ # 対話モード（REPL）で実行
$ qiitan
Welcome to Qiitan Script Interactive Mode! (To exit just type `bye`)
> foo := import("fmt")
LOG: {sprintf: <user-function>, __module_name__: "fmt", print: <user-function>, printf: <user-function>, println: <user-function>}
> foo.println("Hello World!")
Hello World!
LOG: <undefined> object returned. Perhaps the statement has no return or nothing is assigned to it.
>
> bye
Nice chatting with you. Thank you! Bye-bye~.
$
```

対話モードでは、出力した内容および代入した内容が適宜表示されます。

`foo := import("fmt")` で、`foo` 変数に `fmt` モジュールを代入していますが、その `fmt` モジュールには `print()`, `printf()`, `println()`, `sprintf()` が定義されていることが確認できます。

```go
LOG: {__module_name__: "fmt", print: <user-function>, printf: <user-function>, println: <user-function>, sprintf: <user-function>}
```

`foo.println("Hello World!")` で出力していますが、続く `LOG: <undefined> object returned` は、`fmt.println()` に戻り値がない（型がない値が返ってきた）ことを表しています。

### プリ・コンパイルと実行

キーたん語は、なんちゃってプリ・コンパイル型の言語です。

`qiitan` スクリプトの実行だけでなく、その中間ファイル（バイトコード）を出力することができます。
この中間ファイルも `qiitan` インタプリタで実行することができ、プリ・コンパイルが不要なぶんだけ速く実行されます。

```shellsession
$ qiitan -compile ./helloworld.qiitan
compiled: ./helloworld.qtn

$ qiitan ./helloworld.qtn
Hello World!
```

## Install

- [Homebrew](https://brew.sh/index_ja) (**macOS:** x86_64/Intel/AMD64, M1/ARM64, **Linux & Windows WSL2:** x86_64, ARM64, ARM v6)

    ```bash
    brew install qithub-bot/apps/qiitan
    ```
- 単体バイナリ（Windows, macOS, Linux）
    - [Releases ページ](https://github.com/Qithub-BOT/Qiitan-go/releases/latest)から該当 OS + アーキテクチャをダウンロード。

---

## Statuses

このリポジトリでは、以下のような最低限のセキュリティ対策が施されています。

[![go1.14+](https://github.com/Qithub-BOT/Qiitan-go/actions/workflows/go-versions.yml/badge.svg)](https://github.com/Qithub-BOT/Qiitan-go/actions/workflows/go-versions.yml "Unit Tests on Go 1.16, 17 and latest")
[![Platform Tests](https://github.com/Qithub-BOT/Qiitan-go/actions/workflows/platform-test.yml/badge.svg)](https://github.com/Qithub-BOT/Qiitan-go/actions/workflows/platform-test.yml "Test on Win, macOS, Linux")
[![golangci-lint](https://github.com/Qithub-BOT/Qiitan-go/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Qithub-BOT/Qiitan-go/actions/workflows/golangci-lint.yml "Static Analysis")
[![CodeQL](https://github.com/Qithub-BOT/Qiitan-go/actions/workflows/codeQL-analysis.yml/badge.svg)](https://github.com/Qithub-BOT/Qiitan-go/actions/workflows/codeQL-analysis.yml "Vulnerability Scan")
[![codecov](https://codecov.io/gh/Qithub-BOT/Qiitan-go/branch/main/graph/badge.svg?token=uW30s2bK8M)](https://codecov.io/gh/Qithub-BOT/Qiitan-go "Code Coverage")
[![Go Report Card](https://goreportcard.com/badge/github.com/Qithub-BOT/Qiitan-go)](https://goreportcard.com/report/github.com/Qithub-BOT/Qiitan-go "Code Quality")

## 言語開発（コントリビュート）

[![go1.16+](https://img.shields.io/badge/Go-1.16+-blue?logo=go)](https://github.com/Qithub-BOT/Qiitan-go/actions/workflows/go-versions.yml "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/Qithub-BOT/Qiitan-go.svg)](https://pkg.go.dev/github.com/Qithub-BOT/Qiitan-go)

- `qiitan` コマンド（インタプリタ）自体は Go 言語で開発されています。
- スタンドアロン（単体）のバイナリとして動作するように設計されています。

<!--
- キーたん語（`qiitan` 語）は [Tengo 言語](https://github.com/d5/tengo)を独自のモジュールで拡張したものです。
- `qiitan` スクリプトは `tengo` スクリプトの上位互換です。Tengo 言語で書かれた `tengo` スクリプトは `qiitan` インタープリタで実行できますが、`qiitan` スクリプトは Tengo のインタープリタでは実行できません。
-->

---

## License

- [MIT](https://github.com/Qithub-BOT/Qiitan-go/LICENSE) License. Copyright (c) [The Qiitan-go Contributors](https://github.com/Qithub-BOT/Qiitan-go/graphs/contributors).
