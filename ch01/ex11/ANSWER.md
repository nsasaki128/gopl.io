#練習問題の回答と理由
## 回答と理由:長い引数リスト
### 回答
長い引数リストでも実施される。
### 理由
alexaのトップ25サイトで実施したところ問題なく動いた。
```
0.14s   11263 http://Google.com
0.17s   12766 http://Google.com.br
0.18s   12826 http://Google.de
0.18s   12928 http://Google.co.in
0.19s    3566 http://Live.com
0.51s   48021 http://Weibo.com
0.53s   11153 http://Google.co.jp
0.62s   12772 http://Google.co.uk
0.63s  227070 http://Tmall.com
0.68s  150017 http://360.cn
0.70s  602562 http://Sina.com.cn
0.71s   50242 http://Twitter.com
0.80s       0 http://Instagram.com
0.94s  242967 http://Qq.com
1.09s  172696 http://Reddit.com
1.22s  239192 http://Taobao.com
1.25s   86472 http://Wikipedia.org
1.27s  315322 http://Facebook.com
1.32s  507790 http://Youtube.com
1.36s  517031 http://Yahoo.com
1.41s  144813 http://Jd.com
1.59s  195716 http://Sohu.com
1.89s  475535 http://Amazon.com
2.56s    6996 http://Vk.com
10.14s      81 http://Baidu.com
10.14s elapsed
```
### 確認方法
```bash
./test.bash
```

## 回答と理由:応答しないウェブサイト
### 回答
あるウェブサイトが応答しない場合は待ち続けると思われる。
### 理由
応答しないサーバー(server/main.go)を作成し、下記の方法で確認したところ、処理が進まなかったため。
#### 確認方法
```bash
./test2.bash
```
