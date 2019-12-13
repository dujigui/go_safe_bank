package main

import (
	"fmt"
	"sync"
	"time"
)

func main3() {
	now := time.Now()
	var wg sync.WaitGroup
	b := Bank3{
		dc: make(chan int),
		bc: make(chan int),
	}
	b.init()
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

type Bank3 struct {
	balance int
	dc      chan int
	bc      chan int
}

func (b *Bank3) init() {
	go func() {
		for {
			select {
			case amount := <-b.dc:
				b.balance += amount
			case b.bc <- b.balance:
			}
		}
	}()
}

func (b *Bank3) Deposit(amount int) {
	b.dc <- amount
}

func (b *Bank3) Balance() int {
	return <-b.bc
}
