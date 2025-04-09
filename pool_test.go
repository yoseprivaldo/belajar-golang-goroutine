package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T){
	pool := sync.Pool{
        New: func() any {
            return "New"
        },
    }
	group := sync.WaitGroup{}

	pool.Put("Yosep")
	pool.Put("Rivaldo")
	pool.Put("Silaban")

	for i := 0; i< 10; i++ {
		group.Add(1)
		go func ()  {
			defer group.Done()
            data := pool.Get() // Ambil data dari pool
            fmt.Println(data)  // Cetak data
            // pool.Put(data) 
		}()
	}

	group.Wait()
	fmt.Println("Selesai")
}