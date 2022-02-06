# スクリプト・ファイル

## ファイル・フォーマット

- `qiitan` スクリプトはテキストで記載することができます。お好きなテキスト・エディタをお使いください。
    - `qiitan` スクリプトのシンタックス・ハイライトの VSCode 用拡張機能をリポジトリに同梱しています。手動で VSCode の `.vscode/extensions` に `vscode-qiitan-go` ディレクトリをコピーして利用ください。VSCode 再起動後に使えるようになります。
        - [.vscode/extensions](https://github.com/Qithub-BOT/Qiitan-go/tree/main/.vscode/extensions) @ GitHub

### 文字コード

- `qiitan` スクリプトは UTF-8, BOM なしで書かれている必要があります。

### 改行コード

- 改行コードは `LF` (`0x0A`,`\n`) です。

### 拡張子

- `qiitan` で使える拡張子は `.qiitan` および `.tengo` です。（`.qiitan` 推奨）
  - `.tengo` 拡張子は、親言語である [Tengo 言語](https://github.com/d5/tengo)の拡張子です。キーたん語は Tengo 言語の上位互換であるため、`.tengo` ファイルも実行できますが、`tengo` では拡張子 `.qiitan` のファイルは実行できません。
- 特殊な拡張子: `.qtn`
  - `.qtn` 拡張子は、`qiitan` スクリプトの中間ファイルの拡張子です。複数の `qiitan` スクリプト・ファイルで構成されたアプリを、単体のファイルとして配布したい場合に出力します。
  - `.qtn` ファイルの実行にも `qiitan` インタプリタ（`qiitan` コマンド）は必要です。
  - `.qtn` ファイルはバイナリ・データです。テキスト・エディタなどでは編集できません。

## ファイルのアクセス権

- `0644`: 一般的なアクセス権の設定（`cron` などでも実行させたい人向け）
    ```shellsession
    $ ls -l ./hello.qiitan
    -rw-r--r--  1 keinos  staff  49  1 12 14:29 ./hello.qiitan
    ```
- `0400`: 心配性な人向けの設定
    ```shellsession
    $ ls -l ./hello.qiitan
    -r--------  1 keinos  staff  49  1 12 14:29 ./hello.qiitan
    ```
- スクリプト・ファイルの文頭に [`shebang`](https://ja.wikipedia.org/wiki/%E3%82%B7%E3%83%90%E3%83%B3_(Unix)) を付けて、ファイル単体で実行可能にする場合は、「実行」権限が必要になります。
    ```shellsession
    $ hello
    Hello, world!

    $ type hello
    hello is /usr/local/bin/hello

    $ cat /usr/local/bin/hello
    #!/usr/bin/env qiitan

    fmt := import("fmt")
    fmt.println("Hello, world!")

    $ ls -l /usr/local/bin/hello
    -rwxr-xr-x  1 keinos  staff  49  7 12  2021 /usr/local/bin/hello
    ```

## ファイルの実行

`qiitan` スクリプトの実行は、`qiitan` インタプリタ（`qiitan` コマンド）を通して実行します。

```bash
qiitan [options] [script file | byte code file] [script arg]
```

```shellsession
$ qiitan ./hello_world.qiitan
Hello, Qiitan!

$ qiitan ./say_hello_to.qiitan world!
Hello, world!
```

また、バイト・コードに変換済み（中間ファイルにコンパイル済み）の場合も同じ構文です。

```shellsession
$ # 中間ファイルにコンパイル
$ qiitan --compile ./hello_world.qiitan
** snip **

$ # 中間ファイルの実行
$ qiitan ./hello_world.qtn
Hello, Qiitan!
```
