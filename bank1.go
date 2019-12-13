package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	now := time.Now()
	var wg sync.WaitGroup
	b := Bank1{}
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			b.Deposit(100)
			b.Balance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("Final: %d Time: %s\n", b.Balance(), time.Now().Sub(now))
}

type Bank1 struct {
	balance int
}

func (b *Bank1) Deposit(amount int) {
	b.balance += amount
}

func (b *Bank1) Balance() int {
	return b.balance
}
