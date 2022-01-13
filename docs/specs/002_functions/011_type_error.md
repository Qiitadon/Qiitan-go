# `error` type（エラー型と値）

キーたん語では、エラーは `error` 型の値で示すことができます。

## エラーの作成

エラーは、`error()` 関数（組込関数）を使い、戻り値を変数に代入して作成します。

```golang
// string 型の値を持ったエラーの作成
myErr := error("oops") // myErr は error 型の変数
```

```golang
// int 型の値を持ったエラーの作成
errNum := error(1+2+3) // errNum は error 型の変数
```

## エラー内容の取り出し

エラーから、エラーの内容を取り出すには `.value` セレクタを参照します。セレクタは、オブジェクトのプロパティやフィールドと同じ意味です。

```golang
myErr := error("oops")
errVal := myErr.value // errVal = "oops"
```

## エラー型の確認

変数がエラーであるかの確認は `is_error()` 関数（組込関数）を使います。

```golang
err1 := error("oops")
err2 := error(1+2+3)

if is_error(err1) {      // 'is_error' 関数で型の確認
  err_val := err1.value  // 値の参照（err_val = "oops"）
}

if is_error(err2) {
  err_val := err2.value  // 値の参照（err_val = 6）
}
```

---

> [目次に戻る ↩️](../../)
