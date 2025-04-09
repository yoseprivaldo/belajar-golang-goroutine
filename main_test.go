package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func worker(name string, delay time.Duration, result chan<- string) {
	time.Sleep(delay)
	result <- fmt.Sprintf("%s selesai setelah %v", name, delay)
}

func runWorkersWithTimout(timeout time.Duration) string {
	result := make(chan string)

	for i := 1; i <= 3; i++ {
		str := "Worker" + strconv.Itoa(i)
		go worker(str, 5*time.Second, result)
	}

	select {
	case res := <-result:
		return "✅ " + res
	case <-time.After(timeout):
		return "⏰ Timeout"
	}
}

func TestRunWorkersWithTimeout_Success(t *testing.T) {
	result := runWorkersWithTimout(6 * time.Second)

	if !strings.Contains(result, "Worker 2") {
		t.Errorf("Expected Worker 2 to finish first, got: %s", result)
	}
}

func TestRunWorkersWithTimeout_Timeout(t *testing.T) {
	result := runWorkersWithTimout(1 * time.Second)

	if !strings.Contains(result, "Timeout") {
		t.Errorf("Expected timeout, got: %s", result)
	}
}