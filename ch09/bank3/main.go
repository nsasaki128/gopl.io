package main

import "sync"

var (
	mu      sync.Mutex //balanceを保護する
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
