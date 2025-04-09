package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <- timer.C

	fmt.Println(time)
}

func TestAfter(t *testing.T){
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())
	time := <- channel

	fmt.Println(time)
}

func TestAfterFunction(t *testing.T){
	group:= sync.WaitGroup{}

	group.Add(1)

	fmt.Println("Dijalankan pertama kali")
	defer fmt.Println("Dijalankan terakhir")

	time.AfterFunc(5*time.Second, func ()  {
		fmt.Println("Execute after 5 second")
		group.Done()
	})

	group.Wait()
}

func TestTimerNew(t *testing.T) {
	ch := make(chan string)

	go func ()  {
		time.Sleep(2 * time.Second)
		ch <-"Response from service"		
	}()

	select {
	case res:= <-ch:
		fmt.Println("Dapet response: ", res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout! Gak dapat response dalam 3 detik")
	}
}