#!/bin/sh
# =============================================================================
#  DOCS/INDEX.MD ジェネレーター
# =============================================================================
#  このシェル・スクリプトは docs/index.md を更新（作成）します。直接実行するか、
#  リポジトリのルートから `go generate ./...` を実行しても自動生成されます。
#
#  - 要 `gh-md-toc` コマンド: https://github.com/ekalinin/github-markdown-toc.go
#    go install "github.com/ekalinin/github-markdown-toc.go/cmd/gh-md-toc@latest"

# 定数
NAME_FILE_INDEX="index.md"
PATH_DIR_RETURN="$(cd "$(pwd)" && pwd)"
PATH_DIR_SCRIPT="$(cd "$(dirname "${0}")" && pwd)"
PATH_DIR_REPO_ROOT="$(cd "${PATH_DIR_SCRIPT}/../.." && pwd)"
PATH_DIR_DOCS="${PATH_DIR_REPO_ROOT}/docs"
PATH_FILE_INDEX="${PATH_DIR_DOCS}/index.md"
LF=$(printf '\n_');LF=${LF%_}
HR="---${LF}"

# 作業ディレクトリの移動とリカバリ
deferChDir() {
    cd "${PATH_DIR_RETURN}" || echo >&2 "failed to move back: ${PATH_DIR_RETURN}"
}
trap deferChDir 0

cd "${PATH_DIR_DOCS}" || {
    echo >&2 "failed to move: ${PATH_DIR_DOCS}"
    exit 1
}

# 自動生成である旨のヘッダ出力
printf "<!-- Code generated using /gen/gen_readme.sh; DO NOT EDIT. -->\n" >"${PATH_FILE_INDEX}"

# ./modules_tengo/* 下の Markdown の目次一覧を取得
# shellcheck disable=SC2046
modules_tengo="$(gh-md-toc --depth=2 --hide-footer --hide-header --serial $(find ./modules_tengo -name '*.md' | sort) | sed '/^$/d')"

# ./modules_qiitan/* 下の Markdown の目次一覧を取得
# shellcheck disable=SC2046
modules_qiitan="$(gh-md-toc --depth=2 --hide-footer --hide-header --serial $(find ./modules_qiitan -name '*.md' | sort) | sed '/^$/d')"

{
    # タイトル追記
    printf "# Qiitan スクリプトの言語仕様\n\n"
    printf "- [Qiitan-go リポジトリ](https://github.com/Qithub-BOT/Qiitan-go) @ GitHub\n\n"
    printf "## 目次\n\n"

    # ./specs/* 下の Markdown の目次一覧を追記
    # shellcheck disable=SC2046
    gh-md-toc --depth=2 --hide-footer --hide-header --serial $(find ./specs -name '*.md' | sort)

    # ./modules_tengo/* 下の Markdown の目次一覧を追記
    printf "## モジュール\n\n"
    printf "### Tengo 互換モジュール\n\n"
    printf "%s\n\n" "${modules_tengo}"
    printf "### qiitan 独自モジュール\n\n"
    printf "%s\n\n" "${modules_qiitan}"

    # フッターの追記
    echo "${HR}"
    echo '- View the repository: [Qiitan-go](https://github.com/Qithub-BOT/Qiitan-go) @ GitHub'
    echo '- Table of contents created by [gh-md-toc](https://github.com/ekalinin/github-markdown-toc.go)' @ GitHub

} >>"${PATH_FILE_INDEX}"
