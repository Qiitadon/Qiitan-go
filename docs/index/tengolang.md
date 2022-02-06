# Tengo 言語

[Tengo](https://github.com/d5/tengo) 言語は、キーたん語の元になっているスクリプト言語です。

- [The Tengo Language](https://github.com/d5/tengo) @ GitHub

「キーたん語」は、`Tengo` 言語の標準[モジュール](#モジュール)に加え、独自のモジュールを同梱して拡張した Tengo 言語の上位互換言語です。

```bash
qiitan 標準モジュール = tengo 標準モジュール + qiitan 独自拡張モジュール
```

そのため、基本的に `tengo` 用に書かれたスクリプトも `qiitan` インタプリタで実行できます。

```shellsession
$ tengo ./hello_world.tengo
Hello, world!

$ qiitan ./hello_world.tengo
Hello, world!
```

逆に、キーたん語のスクリプトは基本的に `tengo` インタプリタでは実行できません。`.qiitan` 拡張子を解釈できないためです。

`tengo` 互換（キーたん語による拡張モジュールを使わない）で書かれている場合は、拡張子を変えることで実行できます。

```shellsession
$ tengo ./hello_world.qiitan
unexpected EOF

$ mv ./hello_world.qiitan ./hello_world.tengo
$ tengo ./hello_world.tengo
Hello, world!
```

> [目次に戻る ↩️](../)
