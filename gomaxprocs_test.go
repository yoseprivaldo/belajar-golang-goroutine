package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// * untuk mengetahui berapa jumlah thread, bisa menggunakan GOMAXPROCS
// * secara default jumlah thread di GoLang itu sebanyak jumlah CPU di komputer kita
// * kita juga bisa melihat jumlah CPU kita dengan menggunakan function runtime.NumCpu()

func TestGomaxProcs(t *testing.T){
	group := sync.WaitGroup{}

	totalCpu := runtime.NumCPU()
	fmt.Println("CPU", totalCpu)

	// * mengambil jumlah default thread
	// totalThread := runtime.GOMAXPROCS(-1)
	// fmt.Println("Thread", totalThread)

	// * merubah jumlah thread (jarang digunakan)
	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThread)


	group.Wait()
	
	totalGorouiten := runtime.NumGoroutine()
	fmt.Println("Goroutine", totalGorouiten)
}  