![](https://img.shields.io/badge/Tengo%20Compatible-Yes-blue "Tengo compatible module")
![](https://img.shields.io/badge/Module%20type-Go-blue "Module made in Go")

# fmt

`fmt` モジュールは、「標準出力」に出力する関数群です。他の言語でいう `Print` 系の関数です。

```go
fmt := import("fmt")
```

```go
fmt := import("fmt")

fmt.println("Hello, world!")

// Output:
// Hello, world!
```

- [このモジュールの定義ソース](https://github.com/d5/tengo/blob/master/stdlib/fmt.go) @ github.com/d5/tengo

---

## print

```go
print(args...)
// args: string type (variadic)
```

`print` は、`args` の文字列を連結して標準出力に出力します。

- 末尾に改行コードは自動追加しません。
- 引数は可変（`variadic`）で複数引数がある場合、1 つの文字列として結合して出力します。この時、`args` 間にはスペースを入れません。

```go
fmt.print("Hello, world! (single arg)\n")
// Output: Hello, world! (single arg)\n

fmt.print("Hello,", " ", "world!", " ", "(multiple args)", "\n")
// Output: Hello, world! (multiple args)\n

fmt.print("Unlike", "Go's", "fmt.Print", "function,",
    "no spaces are added", "between the operands.", "\n")
// Output: UnlikeGo'sfmt.Printfunction,no spaces are addedbetween the operands.\n
```

> [目次に戻る ↩️](../)

---

## println

```go
println(args...)
// args: string type (variadic)
```

`println` は、`args` の文字列を連結して改行付きで標準出力に出力します。

- OS の改行コードが末尾に自動追加されます。
- 引数は可変（`variadic`）で複数引数がある場合、1 つの文字列として結合して出力します。この時、`args` 間にスペースを入れません。

```go
fmt.println("Hello, world! (single arg)")
// Output: Hello, world! (single arg)\n

fmt.println("Hello,", " ", "world!", " ", "(multiple args)")
// Output: Hello, world! (multiple args)\n

fmt.println("Unlike", "Go's", "fmt.Print", "function,",
    "no spaces are added", "between the operands.")
// Output: UnlikeGo'sfmt.Printfunction,no spaces are addedbetween the operands.\n
```

> [目次に戻る ↩️](../)

---

## printf

```go
printf(format, [args...])
// format: string type
// args: string type
```

`printf` は、`args` を `format` の書式に沿って整形して標準出力に出力する関数です。

- 末尾に改行コードは自動追加しません。
- 引数は可変（`variadic`）で、`format` 内の代替文字（`%v` など）の個数と同数でないといけません。
- `args` が指定されていない場合、`format` の内容が展開せず（`\n` などもパースせず）に出力します。

```go
msg := "world!"

fmt.printf("Hello, %s\n", msg)
// Output: Hello, world!

fmt.printf("Hello, %v\n", msg)
// Output: Hello, "world!"
```

> [目次に戻る ↩️](../)

---

> [目次に戻る ↩️](../)

---
