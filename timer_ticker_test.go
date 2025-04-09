package belajar_golang_goroutine

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

// * ticker adalah representasi kejadian yang berulang
// * ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
// * untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
// * Untuk menghentikan ticker kita bisa menggunakan Ticker.Stop()

func TestTicker(t *testing.T) {
	var counter int64 = 0;
	ticker := time.NewTicker(1 * time.Second)

	for tick := range ticker.C {
		atomic.AddInt64(&counter, 1)
		fmt.Println(tick)

		if counter == 3 {
			ticker.Stop()
			break;
		}
	}
}

func TestTick(t *testing.T) {
	var counter int64 = 0;
	channel := time.Tick(5 * time.Second)

	for tick := range channel{
		atomic.AddInt64(&counter, 1)
		fmt.Println(tick)		
	}

}