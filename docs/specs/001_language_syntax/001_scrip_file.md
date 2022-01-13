# Qiitan スクリプト・ファイル

## ファイル・フォーマット

- `qiitan` スクリプトはテキストで記載することができます。お好きなテキスト・エディタをお使いください。
    - `qiitan` スクリプトのシンタックス・ハイライトの VSCode 用拡張機能をリポジトリに同梱しています。手動で VSCode の `.vscode/extensions` に `vscode-qiitan-go` ディレクトリをコピーして利用ください。VSCode 再起動後に使えるようになります。
        - [.vscode/extensions](https://github.com/Qithub-BOT/Qiitan-go/tree/main/.vscode/extensions) @ GitHub

## 文字コード

- `qiitan` スクリプトは UTF-8, BOM なしで書かれている必要があります。

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

