# Documents for Qiitan Script Spec

このディレクトリは、`qiitan` スクリプトを利用するための各種仕様のドキュメント置き場です。

- [ドキュメントを読む](https://Qiitadon.github.io/Qiitan-go/)

typo、言い回しなど、気付いたら何でも PR してください。

ドキュメントを Wiki にしなかったのは、将来的に PR の CI チェックに日本語の Text Linter を導入する予定があるためです。

## index.md について

- `/docs/index.md` は**自動生成される総目次**です。
- `/docs/` 下にある `*.md` ファイルから、各々の見出し行を抽出した一覧になります。
- 抽出処理は `/gen/gen_index/gen_index.sh` のシェル・スクリプトをご覧ください。

## index.md の更新

- リポジトリのルートから `go generate ./...` を実行してください。

## キーたん語の言語拡張

`qiitan` スクリプトの「言語そのもの」を拡張（`qiitan` インタプリタの機能拡張）をしたい場合は、以下のファイルをご覧ください。

- [キーたん語の拡張](../qiitan/README.md)
