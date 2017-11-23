#回答
下記の通り、strings.Joinを用いたバージョンは他と比較して約2倍速い。
```bash
BenchmarkEcho1-4   	 5000000	       262 ns/op
BenchmarkEcho2-4   	 5000000	       262 ns/op
BenchmarkEcho3-4   	10000000	       122 ns/op
```