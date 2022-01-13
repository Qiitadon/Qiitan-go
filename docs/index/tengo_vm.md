# tengo VM（tengo 仮想マシン）

「tengo VM」とは、[`qiitan` インタプリタ](#qiitan-インタプリタ)に内包されている [Tengo の仮想マシン](https://github.com/d5/tengo/blob/master/vm.go)です。バイトコードに変換されたデータをメモリ内の閉鎖空間で実行するために使われます。

仮想マシンや閉鎖空間といっても、たいそうな仕組みではなく、Go 言語の [sync/atomic](https://pkg.go.dev/sync/atomic) 標準ライブラリを使った[アトミック操作](https://ja.wikipedia.org/wiki/%E4%B8%8D%E5%8F%AF%E5%88%86%E6%93%8D%E4%BD%9C)です。

つまり、`qiitan` スクリプトは以下の状態で実行されます。

1. 全操作が完了するまで、他のプロセスはその途中の状態を観測できない。
2. 一部操作が失敗したら、組合せ全体が失敗し、システムの状態は操作を行う前の状態に戻る。

Docker のコンテナなどにも使われている技術ですが、気になる方は以下の動画を参考にしてみてください。`bash` シェルの実行を、外部コマンドとして `exec()` でコールするも、コンテナ上で実行する 20 分程度の動画です。
  - 【英語】[Building a container from scratch in Go - Liz Rice (Microscaling Systems)](https://youtu.be/Utf-A4rODH8) | Container Camp @ Youtube

> [目次に戻る ↩️](../)
