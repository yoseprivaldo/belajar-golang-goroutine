package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// * terjadi race condition
func TestRaceCondition(t *testing.T){
	var x = 0;
	for i := 1; i <=1000; i++ {
		go func ()  {
			for j := 1; j <= 100; j++ {
				x = x +1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", x)
}

// * solving race condition dengan mutex
func TestSolveRaceConditionWithMutex(t *testing.T) {
	var x =0
	var mutex sync.Mutex
	var group sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func ()  {
			defer group.Done()
			for i := 1; i <= 100; i++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}		
		}()
	}
	group.Wait()
	fmt.Println("Counter: ", x)
}


