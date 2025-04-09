package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()

	account.Balance = account.Balance + amount

	account.RWMutex.Unlock()
}

func (account *BankAccount) getBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance;
}

func TestRWMutex(t *testing.T){
	account := BankAccount{}

	for range 100 {
		go func ()  {
			for range 100 {
				account.AddBalance(1)
				fmt.Println(account.getBalance())
			}
			
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance: ", account.getBalance())
}