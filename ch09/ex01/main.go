package main

type withdrawReq struct {
	amount int
	result chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan withdrawReq)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	c := make(chan bool)
	withdraws <- withdrawReq{amount: amount, result: c}
	return <-c
}

func teller() {
	var balance int //balanceはtellerゴルーチンに閉じ込められている
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdraws:
			if balance < w.amount {
				w.result <- false
				continue
			}
			balance -= w.amount
			w.result <- true
		}

	}
}

func init() {
	go teller() //モニターゴルーチンを開始する
}
