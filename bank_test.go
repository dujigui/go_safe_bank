package main

import (
	"sync"
	"testing"
)

func BenchmarkBank1(b *testing.B) {
	var wg sync.WaitGroup
	bank := Bank1{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			bank.Deposit(100)
			bank.Balance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	//fmt.Printf("Bank1 Final: %d\n", bank.Balance())
}


func BenchmarkBank2(b *testing.B) {
	var wg sync.WaitGroup
	bank := Bank2{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			bank.Deposit(100)
			bank.Balance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	//fmt.Printf("Bank2 Final: %d\n", bank.Balance())
}

func BenchmarkBank3(b *testing.B) {
	var wg sync.WaitGroup
	bank := Bank3{
		dc: make(chan int),
		bc: make(chan int),
	}
	bank.init()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			bank.Deposit(100)
			bank.Balance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	//fmt.Printf("Bank3 Final: %d\n", bank.Balance())
}

func BenchmarkBank4(b *testing.B) {
	var wg sync.WaitGroup
	bank := Bank4{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			bank.Deposit(100)
			bank.Balance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	//fmt.Printf("Bank4 Final: %d\n", bank.Balance())
}