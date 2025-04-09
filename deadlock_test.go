package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserAccount struct{
	Mutex sync.RWMutex
	Username string
	Balance int
	TransactionTime time.Time

}

func (user *UserAccount) Lock(){
	user.Mutex.Lock()
}

func (user *UserAccount) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserAccount) Change(amount int) {
	user.Balance = user.Balance + amount
}

// * ini akan terjadi deadlock
func TransactionDeadlock(user1 *UserAccount, user2 *UserAccount, amount int) {
	user1.Lock()
	fmt.Println("Lock user 1", user1.Username)
	user1.Change(-amount)
	
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.Username)
	user2.Change(amount)


	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

// * solusi agar tidak deadlock, lakukan pengurutan penguncian
// * berdasarkan salah satu field bisa jadi tanggal transaksi
func TransactionSolveDeadlock(user1 *UserAccount, user2 *UserAccount, amount int) {
	if user1.TransactionTime.Before(user2.TransactionTime) {
		user1.Lock()
		fmt.Println("Lock user 1", user1.Username)
		user2.Lock()
		fmt.Println("Lock user 2", user1.Username)		
	} else {
		user2.Lock()
		fmt.Println("Lock user 2", user1.Username)		
		user1.Lock()
		fmt.Println("Lock user 1", user2.Username)		
	}

    // * lakuktan transaksi
	user1.Change(-amount)
	user2.Change(amount)
	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T){
	user1 := UserAccount{
		Username: "Yosep",
		Balance: 100000,
		TransactionTime: time.Now(),
	}
	time.Sleep(10 * time.Millisecond)
	user2 := UserAccount{
		Username : "Rivaldo",
		Balance: 100000,
		TransactionTime: time.Now(),
	}

	// go TransactionDeadlock(&user1, &user2, 10000)
	// go TransactionDeadlock(&user2, &user1, 20000)
	go TransactionSolveDeadlock(&user1, &user2, 10000)
	go TransactionSolveDeadlock(&user2, &user1, 20000)

	time.Sleep(2 * time.Second)
	fmt.Println("Balance dari user 1", user1.Balance)
	fmt.Println("Balance dari user 2", user2.Balance)
}
