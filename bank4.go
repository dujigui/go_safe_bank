package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main4() {
	now := time.Now()
	var wg sync.WaitGroup
	b := Bank4{}
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

type Bank4 struct {
	balance int64
}

func (b *Bank4) Deposit(amount int64) {
	atomic.AddInt64(&b.balance, amount)
}

func (b *Bank4) Balance() int64 {
	return atomic.LoadInt64(&b.balance)
}
