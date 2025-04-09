package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0;

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T){
	once:= sync.Once{}
	group := sync.WaitGroup{}
	
	for range 100 {
		group.Add(1)
		go func ()  {
			once.Do(OnlyOnce)
			group.Done()
		}()
	}
	group.Wait();
	fmt.Println(counter)

}