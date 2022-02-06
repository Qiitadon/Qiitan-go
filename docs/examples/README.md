# qiitan スクリプトのサンプル

このディレクトリには qiitan スクリプトの各種サンプルが置いてあります。

    This is a code bloc data
    Do you read this?

```go
// Sample 1 (no output)
fmt := import("fmt")

fmt.println("Hello, world")
```

```go qiitan
// Sample 2 (inline output)
fmt := import("fmt")

fmt.println("Hello, world")
// Output: Hello, world
```

```go qiitan
// Sample 3 (block output)
fmt := import("fmt")

fmt.println("Hello, world")
// Output:
// Hello, world
```
