[[目次](../)]

# キーたん語（qiitan スクリプト）

キーたん語（`Qiitan-go`）は、 [`qiitan` インタプリタ](#qiitan-インタプリタ)を介して実行することができる、中間言語方式の「なんちゃってプリ・コンパイル型」スクリプト言語です。

キーたん語は、[`Tengo`](https://github.com/d5/tengo) 言語（`Tengo v2`）を独自に拡張したスクリプト言語で、この「キーたん語」で書かれたスクリプトを「qiitan スクリプト」と呼びます。

## 「キーたん語」と「qiitan スクリプト」の違い

基本的に同じです。

主に構文や文法などの言語仕様や言語そのものを説明する時は「キーたん語」を使い、キーたん語で書かれた内容やファイルを「qiitan スクリプト」と表しています。こだわりはありませんが、`qiitan` と小文字を使うようにしてください。

## Qiitan と qiitan の使い分け

このドキュメントでは以下のように使い分けています。

- `Qiitan`（大文字）
  - リポジトリ名である `Qiitan-go` を除き、[Qiita](https://qiita.com/) の SNS である [Qiitadon](https://qiitadon.com/) のマスコット・キャラクター「[キーたん](https://github.com/increments/mastodon/blob/qiitadon/app/javascript/images/qiitadon-getting-started.png)」を英字で指す場合に大文字「`Qiitan`」と表します。
- `qiitan`（小文字）
  - 「キーたん語」に関係するものを指す場合は、小文字の `qiitan` が使われます。
  - `qiitan` スクリプトを実行するインタプリタのバイナリ本体（コマンド名）を表す場合も `qiitan` が使われます。

> [目次に戻る ↩️](../)
