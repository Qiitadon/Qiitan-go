<!-- markdownlint-disable MD041 -->
# Qiitan-go<sub><sup><sup>alpha</sup></sup></sub>

[キーたん語（Qiitan-go）](https://github.com/Qithub-BOT/Qiitan-go)は、Qiitan ファンの、Qiitan ファンによる、Qiitan ファンのための「お遊び用スクリプト言語」です。

> Qiitan は [Qiita](https://qiita.com/)™️ の SNS「[Qiitadon](https://qiitadon.com/)<sub><sup><sup>β</sup></sup></sub>」のマスコット・キャラクターです。
>
> - [キーたん（Qiitan）](https://github.com/increments/mastodon/blob/qiitadon/app/javascript/images/qiitadon-getting-started.png) @ Qiitadon
> - `Qiitan-go` は Qiitan のファン・アプリであり、[Qiita](https://qiita.com/)™️ とは一切関係がありません。

---

[WIP](./.qiitask/todo.txt)

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
