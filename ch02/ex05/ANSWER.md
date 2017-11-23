#回答
速度は以下の順となる。

  1. テーブル参照を単一の式で行ったもの(PopCount)
  1. 0になるまで最下位bitのクリアを行った回数を求めるもの(ClearPopCount)
  1. テーブル参照のループを用いたもの(IteratePopCount)
  1. ビットシフトを繰り返し、最下位のビットを参照したもの(ShiftPopCount)


なお、おまけとしてHacker's Delightに記載されていた分割統治法を用いたpopulation countを64bit版に拡張したDivideAndConquerPopCountを試したら最速であった

```bash
BenchmarkPopCount-4                   	2000000000	         0.38 ns/op
BenchmarkIteratePopCount-4            	100000000	        20.7 ns/op
BenchmarkShiftPopCount-4              	20000000	        85.5 ns/op
BenchmarkClearPopCount-4              	100000000	        11.8 ns/op
BenchmarkDivideAndConquerPopCount-4   	2000000000	         0.36 ns/op
```