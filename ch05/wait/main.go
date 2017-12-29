package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// WaitForServer はURLのサーバへ接続を試みます。
// 紙数バックオフを使って一分間試みます。
// 全ての試みが失敗したらエラーを報告します。

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil //成功
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) //紙数バックオフ
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
