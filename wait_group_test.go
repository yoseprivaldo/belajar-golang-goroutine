package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup, counter *int, mutex *sync.Mutex) {
	defer group.Done()

	mutex.Lock()
	*counter++
	fmt.Println("Goroutine finished, remaining", *counter)
	mutex.Unlock()

	time.Sleep(1 * time.Second)

	mutex.Lock()
	*counter--
	fmt.Println("Gorouting finished, remaining: ", *counter)
	mutex.Unlock()
} 

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}
	var counter int;
	var mutex sync.Mutex

	for range 100 {
		group.Add(1)
		go RunAsynchronous(group, &counter, &mutex)
	}

	group.Wait()
	fmt.Println("Complete")
}