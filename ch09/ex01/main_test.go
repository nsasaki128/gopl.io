// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	// Alice
	Deposit(300)
	fmt.Println("=", Balance())

	// Bob
	if !Withdraw(500) {
		t.Errorf("Deposit(300); Withdraw(500) = true, want false")
	}

	//Charlie
	if Withdraw(200) {
		t.Errorf("Deposit(300); Withdraw(200) = false, want true")
	}

	if got, want := Balance(), 100; got != want {
		t.Errorf("Deposit(300); Withdraw(200); Balance() = %d, want %d", got, want)
	}
}
