package main

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} //トークンを獲得
	balance = balance + amount
	<-sema //トークンを解放
}

func Balance() int {
	sema <- struct{}{} //トークンを獲得
	b := balance
	<-sema //トークンを解放
	return b
}
