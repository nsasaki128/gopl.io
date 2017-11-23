#回答
速度は以下の順となる。

  1. テーブル参照を単一の式で行ったもの(PopCount)
  1. テーブル参照のをループを用いたもの(IteratePopCount)
```bash
BenchmarkPopCount-4          	2000000000	         0.40 ns/op
BenchmarkIteratePopCount-4   	100000000	        20.3 ns/op
```