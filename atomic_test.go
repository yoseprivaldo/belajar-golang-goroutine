package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// * solving race condition dengan atomic
// * atomic cocok digunakna untuk tipe data primitif
func TestAtomic(t *testing.T) {
	var x int64 =0
	var group sync.WaitGroup

	for i := 1; i <= 1000; i++ {
		group.Add(1)
		go func ()  {
			defer group.Done()
			for i := 1; i <= 100; i++ {
				atomic.AddInt64(&x, 1)
			}		
		}()
	}
	group.Wait()
	fmt.Println("Counter: ", x)
}


