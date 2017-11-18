#練習問題の回答と理由
## 回答
キャッシュされている。
## 理由
wikipediaのgo言語のページで実行した結果短縮されたため。
go run main.go https://en.wikipedia.org/wiki/Go_\(programming_language\)
```
1.09s  272291 https://en.wikipedia.org/wiki/Go_(programming_language)
1.09s elapsed
0.18s  272291 https://en.wikipedia.org/wiki/Go_(programming_language)
0.18s elapsed
```
上記を複数回実行したが、いずれも2回目の時間は短縮されていた。