# qiitan インタプリタ

`qiitan` [インタプリタ](https://ja.wikipedia.org/wiki/%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF)は、`qiitan` スクリプトの実行もしくは[バイトコード](../#バイトコードとコンパイル)（中間ファイル）に変換するためのコマンド（CLI アプリ）です。Go 言語で書かれており、バイナリ単体で動作します。

```shellsession
$ # qiitan スクリプトを qiitan インタプリタで実行する例
$ qiitan ./hello_world.qiitan
Hello world!
```

`qiitan` インタプリタは、一旦バイトコードを作成してから [Tengo の仮想マシン](#tengo-vmtengo-仮想マシン)上でバイトコードを実行します。そのため、同じスクリプトを複数回実行する場合は、バイトコードに変換したものを利用することをオススメします。

> [目次に戻る ↩️](../)
