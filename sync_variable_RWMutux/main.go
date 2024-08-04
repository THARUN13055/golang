package main

import (
	"errors"
	"fmt"
	"sync"
)

/*
here we are adding 3 func like balance , deposit, withdraw , main func
balance how many time we can able to readonly, withdraw and deposit we need to use only lock and unloack
*/

type Amount struct {
	balance int
	sync.RWMutex
}

func (a *Amount) Balance() int {
	a.RLock()
	defer a.RUnlock()
	return a.balance
}

func (a *Amount) deposit(getting int) {
	a.Lock()
	a.balance = a.balance + getting
	a.Unlock()
}

func (a *Amount) Withdraw(taking int) error {
	a.Lock()
	defer a.Unlock()
	if a.balance >= taking {
		a.balance = a.balance - taking
		return nil
	} else {
		return errors.New("Amount is in sufficient")
	}
}

func main() {
	value := Amount{balance: 1000}
	done := make(chan struct{}, 2)
	go func() {
		defer func() { done <- struct{}{} }()
		err := value.Withdraw(200)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("successfully withdraw 200")
		}

	}()
	go func() {
		defer func() { done <- struct{}{} }()
		err := value.Withdraw(400)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("successfully withdraw 400")
		}
	}()
	go func() {
		defer func() { done <- struct{}{} }()
		value.deposit(500)
	}()
	<-done
	<-done
	<-done
	fmt.Println(value.balance)
}
