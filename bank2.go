package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {
	now := time.Now()
	var wg sync.WaitGroup
	b := Bank2{}
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

type Bank2 struct {
	balance int
	mu      sync.RWMutex
}

func (b *Bank2) Deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}

func (b *Bank2) Balance() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.balance
}
