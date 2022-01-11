# コード & ドキュメント・ジェネレーター

このディレクトリは各種コードやドキュメントを自動生成・更新するためのものです。

## 更新方法

リポジトリのルートから `go generate ./...` コマンドを実行します。

また、個別に実行したい場合は `$ go generate ./modules/modules.go` のように、`go:generate` コメント行の記載のある Go ファイルを指定します。

- 要 `gh-md-toc`
	```bash
	go install "github.com/ekalinin/github-markdown-toc.go/cmd/gh-md-toc@latest"
	```

## 更新の仕組み

このリポジトリ（`qiitan-go`）では、ドキュメントやコードの自動生成に Go の `generate` コマンド利用しています。

`go generate` コマンドは、`*.go` ファイルのコメント行に `//go:generate <コマンド>` があった場合、その `<コマンド>` を実行します。

更新方法は、リポジトリのルート・ディレクトリから `$ go generate ./...` を実行すると、全ての自動生成が実行されます。

## 自動生成されるファイル一覧

- `/modules/ModulesInScript.go`
  - この Go ファイルは Tengo 言語で書かれたスクリプトを標準モジュールとして使えるようにするためのものです。
  - 呼出： `/modules/modules.go` の `go:generate` 文
  - 作成： `/gen/gen_scrmods/main.go` を実行することにより作成されます。
  - タイミング： スクリプト・モジュール（`/modules/scrmods/scrmod_*.tengo`）を追加・変更を行った場合に更新される必要があります。

- `/docs/README.md`
  - このファイルは `/docs/spec`s 以下にある Markdown ファイルの総目次です。
  - 呼出： `/gen/gen_scrmods/main.go` の `go:generate` 文
  - 作成： `/gen/gen_readme/gen_readme.sh` を実行することにより作成されます。
  - タイミング： `/docs/specs` 以下のファイルが追加・変更された場合に更新される必要があります。
  - 注意： Markdown の生成に GitHub API を利用しています。API の制限枠があるので、5 回/時 以上更新する場合には注意が必要です。詳しくは

## 注意（GitHub API の利用制限）

デバッグ中など、1 時間で 5 回以上 `go generate` を実行した場合、"`API rate limit exceeded for ...`" エラー発生時する場合があります。その場合は GitHub API のアクセストークンが必要です。

これは、Markdown の目次作成に使われる `gh-md-toc` コマンドが、GitHub 互換を持たせるために Markdown の解析に GitHub API を使用しているからです。

アクセストークンなしの場合、60 リクエスト／時という GitHub API の制限に触れてしまい、"API rate limit exceeded for ..." エラーが発生します。

アクセストークンは [GitHub 上でパーソナル・アクセストークンを発行](https://github.com/settings/tokens)でき、実行している環境の環境変数 `GH_TOC_TOKEN` に指定しておきます。

発行時のトークンのスコープ（アクセス権限）は、特に指定なし（公開情報の読み取りのみ）で十分です。

コンテナで開発している（`.devcontainer` を利用している）場合は、ローカルの環境変数 `GH_TOC_TOKEN` に設定しておくと、コンテナにも反映されます。

- [/.devcontainer/devcontainer.env](../.devcontainer/devcontainer.env)
