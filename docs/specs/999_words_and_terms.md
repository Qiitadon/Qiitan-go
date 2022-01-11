[[目次](../README.md#用語集)]

---

# 用語集

「キーたん語」の言語仕様もしくは「qiitan スクリプト」の構文など、仕様の中で使われる用語の定義と説明です。（順不順）

---

## キーたん語（qiitan スクリプト）

キーたん語（`Qiitan-go`）は、 [`qiitan` インタプリタ](#qiitan-インタプリタ)を介して実行することができる、中間言語方式の「なんちゃってプリ・コンパイル型」スクリプト言語です。

キーたん語は、[`Tengo`](https://github.com/d5/tengo) 言語（`Tengo v2`）を独自に拡張したスクリプト言語で、この「キーたん語」で書かれたスクリプトを「qiitan スクリプト」と呼びます。

### 「キーたん語」と「qiitan スクリプト」の違い

基本的に同じです。

主に構文や文法などの言語仕様や言語そのものを説明する時は「キーたん語」を使い、キーたん語で書かれた内容やファイルを「qiitan スクリプト」と表しています。こだわりはありませんが、`qiitan` と小文字を使うようにしてください。

### Qiitan と qiitan の使い分け

このドキュメントでは以下のように使い分けています。

- `Qiitan`（大文字）
  - リポジトリ名である `Qiitan-go` を除き、[Qiita](https://qiita.com/) の SNS である [Qiitadon](https://qiitadon.com/) のマスコット・キャラクター「[キーたん](https://github.com/increments/mastodon/blob/qiitadon/app/javascript/images/qiitadon-getting-started.png)」を英字で指す場合に大文字「`Qiitan`」と表します。
- `qiitan`（小文字）
  - 「キーたん語」に関係するものを指す場合は、小文字の `qiitan` が使われます。
  - `qiitan` スクリプトを実行するインタプリタのバイナリ本体（コマンド名）を表す場合も `qiitan` が使われます。

> [目次に戻る ↩️](../)

---

## qiitan インタプリタ

`qiitan` [インタプリタ](https://ja.wikipedia.org/wiki/%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF)は、`qiitan` スクリプトの実行もしくは[バイトコード](#バイトコードとコンパイル)（中間ファイル）に変換するためのコマンド（CLI アプリ）です。Go 言語で書かれており、バイナリ単体で動作します。

```shellsession
$ # qiitan スクリプトを qiitan インタプリタで実行する例
$ qiitan ./hello_world.qiitan
Hello world!
```

`qiitan` インタプリタは、一旦バイトコードを作成してから [Tengo の仮想マシン](#tengo-vmtengo-仮想マシン)上でバイトコードを実行します。そのため、同じスクリプトを複数回実行する場合は、バイトコードに変換したものを利用することをオススメします。

> [目次に戻る ↩️](../)

---

## Tengo 言語

[Tengo](https://github.com/d5/tengo) 言語は、キーたん語の元になっているスクリプト言語です。

- [The Tengo Language](https://github.com/d5/tengo) @ GitHub

キーたん語は、`Tengo` v2 の拡張言語になるため基本的に `tengo` 用に書かれたスクリプトも `qiitan` インタプリタで実行できます。

```shellsession
$ tengo ./hello_world.tengo
Hello, world!

$ qiitan ./hello_world.tengo
Hello, world!
```

逆に、キーたん語のスクリプトは基本的に `tengo` インタプリタでは実行できません。`.qiitan` 拡張子を解釈できないためです。

`tengo` 互換（キーたん語による拡張[モジュール](#モジュール)を使わない）で書かれている場合は、拡張子を変えることで実行できます。

```shellsession
$ tengo ./hello_world.qiitan
unexpected EOF

$ mv ./hello_world.qiitan ./hello_world.tengo
$ tengo ./hello_world.tengo
Hello, world!
```

> [目次に戻る ↩️](../)

---

## モジュール

「モジュール」とは、スクリプト内の `import()` で利用可能な「関数群」です。他の言語でいうパッケージに近いものです。

モジュールには「標準モジュール」と「外部モジュール」の 2 種類があります。

### 標準モジュール

「標準モジュール」とは、`qiitan` インタプリタに同梱されているモジュールで、`import()` で「モジュール名の指定」だけで使えるモジュールを言います。

スクリプト内で `hoge := import(<モジュール名>)` と指定すると、`hoge` 変数にモジュールが代入され、モジュール内の関数を `hoge.<モジュール内の関数名>()` のように利用できます。

```go
hoge := import("fmt") // fmt 標準モジュールを hoge に割り当て

hoge.print("fuga") // fmt モジュール内の（モジュール内で export されている）print 関数を利用

// Note: 通常は `fmt := import("fmt")` のように変数名をモジュール名に合わせるのが慣習ですが、別名でも動作します。
```

「キーたん語」は `Tengo` 言語の標準モジュールに加え、独自のモジュールを同梱して拡張したものです。
（`qiitan` 標準モジュール = `tengo` 標準モジュール + `qiitan` 拡張モジュール）

なお、`qiitan` スクリプトを書く上で特に気にする必要はありませんが、`qiitan` の標準モジュールにも、内部的には 2 種類あります。Go 言語で書かれたモジュールと、`Tengo` 言語で書かれたモジュールの 2 種類です。

どちらもバイトコードに変換されてから実行されるので、同等に機能し、コンパイル後の実行速度にも違いはありません。しかし、`qiitan` インタプリタに標準モジュールを追加したい場合（「キーたん語」を、さらに拡張したい場合）には、どちらの言語で開発するかを意識する必要があります。

### 外部モジュール

外部モジュールは、外部スクリプト・ファイルと同等です。

「`Tengo` 言語」および「キーたん語」で書かれたスクリプト・ファイルを、スクリプト内で `import()` で呼び出して使えるモジュールのことを「外部モジュール」と言います。

```go
hoge := import("./fuga") // ./fuga.qiitan もしくは ./fuga.tengo のファイルを読み込み hoge に割り当てる

hoge.piyo("mogera") // fuga モジュール内で export されている piyo 関数を利用
```

> [目次に戻る ↩️](../)

---

## バイトコードとコンパイル

「バイトコード」とは、`qiitan` インタプリタにより作成される中間データのことです。また、キーたん語における「コンパイル」とは、この中間データに変換することを指します。

キーたん語は、なんちゃってプリ・コンパイル型の言語です。`qiitan` インタプリタは、スクリプトを一旦バイトコードに変換してから `tengo` の仮想マシン上でバイトコードを実行します。

このバイトコードはファイル（`*.qtn`）として出力することができ、同じく `qiitan` インタプリタで実行することができます。この場合、バイトコードへのコンパイル（変換）が不要なぶんだけ実行が速くなります。

```shellsession
$ # 通常の実行（コンパイル & 実行）
$ qiitan ./hello_world.qiitan
Hello World!

$ # コンパイルのみ
$ qiitan --compile ./hello_world.qiitan
out: ./hello_world.qtn

$ # コンパイル済みファイルの実行
$ qiitan ./hello_world.qtn
Hello World!
```

> [目次に戻る ↩️](../)

---

## tengo VM（tengo 仮想マシン）

「tengo VM」とは、[`qiitan` インタプリタ](#qiitan-インタプリタ)に内包されている [Tengo の仮想マシン](https://github.com/d5/tengo/blob/master/vm.go)です。バイトコードに変換されたデータをメモリ内の閉鎖空間で実行するために使われます。

仮想マシンや閉鎖空間といっても、たいそうな仕組みではなく、Go 言語の [sync/atomic](https://pkg.go.dev/sync/atomic) 標準ライブラリを使った[アトミック操作](https://ja.wikipedia.org/wiki/%E4%B8%8D%E5%8F%AF%E5%88%86%E6%93%8D%E4%BD%9C)です。

つまり、`qiitan` スクリプトは以下の状態で実行されます。

1. 全操作が完了するまで、他のプロセスはその途中の状態を観測できない。
2. 一部操作が失敗したら、組合せ全体が失敗し、システムの状態は操作を行う前の状態に戻る。

Docker のコンテナなどにも使われている技術ですが、気になる方は以下の動画を参考にしてみてください。`bash` シェルの実行を、外部コマンドとして `exec()` でコールするも、コンテナ上で実行する 20 分程度の動画です。
  - 【英語】[Building a container from scratch in Go - Liz Rice (Microscaling Systems)](https://youtu.be/Utf-A4rODH8) | Container Camp @ Youtube

> [目次に戻る ↩️](../)

---

## オブジェクト

**キーたん語は[オブジェクト指向型言語](https://ja.wikipedia.org/wiki/%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E6%8C%87%E5%90%91%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0)ではありません**が、全てが「型を持った値」であり、[オブジェクト](https://ja.wikipedia.org/wiki/%E3%82%AA%E3%83%96%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88_(%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0))と呼ばれます。

> [目次に戻る ↩️](../)
