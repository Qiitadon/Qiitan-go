# scrmods

scrmods は `tengo` もしくは `qiitan` スクリプトで書かれた、キーたん語の標準モジュールの実装です。

`qiitan` スクリプトの標準モジュールを Tengo 言語や、キーたん語で書くことができます。

## Tengo 言語モジュールのファイル名

スクリプトによるモジュールはファイル単位で、ファイル名に含まれた名前がモジュール名になります。

```html
scrmod_<モジュール名>.tengo
```

- 例

`scrmod_hello.tengo` の場合、"`import("hello")` で利用できるようになります。

## 必須要件

他のスクリプトからモジュールとして使う関数は、スクリプト内で `export{}` で定義する必要があります。

- 例

```go
export {
    fn: func(...args) {
        text := import("text")
        args = append(args, ",", "how are you?")

        return text.join(args, " ")
    }
}
```

