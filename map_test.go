package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

//* golang memiliki sebuah struct bernama sync.map
//* map ini mirip go lang map, namun yang membedakan, map ini aman untuk menggunakan
//* concurrent menggunakan goroutine

func addToMap(data *sync.Map,value int, group *sync.WaitGroup){
	defer group.Done()
	data.Store(value, value)
}

func TestMapSync(t *testing.T){
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i:= range 100{
		group.Add(1)
		go addToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, "", value)
		return true
	})
}

func TestMapSync2(t *testing.T){
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	addToMap := func (value int, group *sync.WaitGroup)  {
		defer group.Done()
		data.Store(value, value)
	}

	for i:= 0; i< 100; i++{
		group.Add(1)
		go addToMap(i, group)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, "", value)
		return true
	})
}