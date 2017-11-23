#回答
速度は以下の順となる。

  1. テーブル参照を単一の式で行ったもの(PopCount)
  1. テーブル参照のをループを用いたもの(IteratePopCount)
  1. ビットシフトを繰り返し、最下位のビットを参照したもの(ShiftPopCount)


```bash
BenchmarkPopCount-4          	2000000000	         0.37 ns/op
BenchmarkIteratePopCount-4   	100000000	        21.1 ns/op
BenchmarkShiftPopCount-4     	20000000	        87.4 ns/op
```