# フィボナッチ数のサンプル

フィボナッチ数を算出する関数の例です。

複数のプログラム言語で同じ処理をしているソースも置いてあります。書き方がどう違うのかの参考にしてください。

なお、Docker がインストールされていれば、`docker-compose` で各々のソースを実行して簡易的な速度測定をすることができます。

## How to run

```bash
# イメージのビルド
docker-compose build
```
```bash
# Go run (ビルドではない)
docker-compose run --rm go
# PHP8
docker-compose run --rm php
# Python3
docker-compose run --rm python
# Pure Qiitan
docker-compose run --rm qiitan
# Qiitan のフィボナッチ・モジュールを使った例
docker-compose run --rm qiitan_std
```

