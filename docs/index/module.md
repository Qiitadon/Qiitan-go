# モジュール

「モジュール」とは、スクリプト内の `import()` で利用可能な「関数群」です。他の言語でいうパッケージに近いものです。

モジュールには「標準モジュール」と「外部モジュール」の 2 種類があります。

---

## 標準モジュール

「標準モジュール」とは、`qiitan` インタプリタに同梱されているモジュールです。標準モジュールは `import()` で「モジュール名の指定」だけで使えます。

具体的には、スクリプト内で `hoge := import(<モジュール名>)` と指定すると、`hoge` 変数にモジュールが代入され、モジュール内の関数を `hoge.<モジュール内の関数名>()` のように利用できます。

注意点として、この時インポート先で利用できるモジュールの関数は、モジュール内で `export` されたもののみになります。

```go
hoge := import("fmt") // fmt 標準モジュールを hoge に割り当て

hoge.println("fuga") // fmt モジュール内の（モジュール内で export されている）println 関数を利用

// Note: 通常は `fmt := import("fmt")` のように変数名をモジュール名に合わせるのが慣習ですが、別名でも動作します。
```

### Go モジュールとスクリプト・モジュール

`qiitan` スクリプトを書く上で特に気にする必要はありませんが、`qiitan` の標準モジュールには、内部的には 2 種類あります。

Go 言語で書かれた「Go モジュール」と、キーたん語もしくは Tengo 言語で書かれた「スクリプト・モジュール」の 2 種類です。

- [標準 Go モジュール](https://github.com/Qithub-BOT/Qiitan-go/tree/main/qiitan/modules/gomods)
- [標準スクリプト・モジュール](https://github.com/Qithub-BOT/Qiitan-go/tree/main/qiitan/modules/scrmods)

これは、**キーたん語の標準モジュールに、新しいモジュールを追加したい**場合に、`qiitan` スクリプトで書いたものでも機能拡張できるということでもあります。つまり、Go 言語に詳しくなくても機能実装できるということです。

逆に、Go モジュールは、Go 言語に慣れている人によって書かれたか、スクリプト・モジュールを高速化をしたい場合に Go に移植されたモジュールです。

Go 言語で書かれたモジュールは、スクリプト・モジュールと比べて圧倒的に速いです。
例えばフィボナッチ数 35 を 5 回求めるのに、`qiitan` スクリプトだけで書いた場合、5 秒かかるのに対し、同じ処理を Go で組み込んだ場合は 0.08 秒まで速くなります。

| Lang | Sec |
| :--: | :--- |
| Qiitan StdMod | 0.08927 (±0.00184) |
| Go run | 0.5119 (±0.0606) |
| PHP8 | 1.4043 (±0.0176) |
| Python3 | 4.6225 (±0.0225) |
| Pure Qiitan | 5.031 (±0.102) |

- [Examples of Fibonacci](https://github.com/Qithub-BOT/Qiitan-go/tree/main/examples/fibonacci) より

しかし、ここでのポイントは、キーたん語でモジュールを書いておけば標準機能として組み込めるということです。

スクリプト・モジュールの API がよくできていれば（理解しやすかったり、動作を変えることなく移植しやすければ）、Go モジュール化される可能性はグッと高くなります。

「こんなモジュールを `qiitan` スクリプトで書いたけど、リポジトリにどうやって組み込んで PR したらいいのかわからない」という場合、[ディスカッション](https://github.com/Qithub-BOT/Qiitan-go/discussions/categories/%E3%81%93%E3%82%93%E3%81%AA%E3%82%93%E4%BD%9C%E3%82%8A%E3%81%BE%E3%81%97%E3%81%9F)にポストしてみてください。

> [目次に戻る ↩️](../)

---

## 外部モジュール

外部モジュールは、外部スクリプト・ファイルと同等です。

「`Tengo` 言語」および「キーたん語」で書かれたスクリプト・ファイルを、スクリプト内で `import()` で呼び出して使えるモジュールのことを「外部モジュール」と言います。

```go
// ./fuga.qiitan もしくは ./fuga.tengo のファイルを読み込み hoge に割り当てる
hoge := import("./fuga")

// fuga モジュール内で export されている piyo 関数を利用
hoge.piyo("mogera")
```

> [目次に戻る ↩️](../)

---
