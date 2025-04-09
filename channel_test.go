package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T){
	channel := make(chan string)

	channel <- "Yosep"

	fmt.Println(<- channel)

	defer close(channel)
}

func TestCreateChannel2(t *testing.T) {
	channel := make(chan string)

	go func ()  {
		time.Sleep(2 * time.Second)
		channel <- "Yosep Rivaldo Silaban"
	}()

	data := <- channel

	fmt.Println(data)
	defer close(channel)
}


// * test channel as paramter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Yosep Rivaldo Silaban"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	fmt.Println(<- channel)

	defer close(channel)
}


// * channel in dan out
//? Menerima channel untuk mengirimkan data ke channel tersebut
func OnlyIn(channel chan<- string) {
	//? jadi channel di block berikut hanya bisa di isi datanya
	// fmt.Println(<- channel)  -> tidak bisa digunakan

	time.Sleep(2 * time.Second)
	channel <- "Yosep Rivaldo Silaban"
}

//? Menerima channel untuk membaca data dari channel tersebut
func OnlyOut(channel <-chan string) {
	//? jadi channel di blok berikut hanya bisa dibaca datanya 
	// channel <- "Rivaldo Silaban" -> tidak bisa

	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	defer close(channel)
}

// * Buffered Channel
func TestBufferedChannel(t *testing.T){
	channel := make(chan string, 3)
	defer close(channel)

	go func ()  {
		channel <- "Yosep"
		channel <- "Rivaldo"
		channel <- "Silaban"
	}()

	go func ()  {
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	} ()


	fmt.Println("Selesai")
}

// * range channel
// * jika channel di kirim terus meneru oleh pengirim
// * jika tidak jelas kapan channel tersebut berhenti menerima data

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func () {
		for i := range 10 { 	
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		go func (data string)  {
			fmt.Println(data)
		}(data)
	}	

	fmt.Println("Done")
}

// * select channel
// * jika kita membuat beberapa channel, dan menjalankan beberapa go routine
// * lalu kita ingin mendapatkan data dari semua channel tersebut

func TestSelectChannel(t *testing.T){
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	// var wg sync.WaitGroup
	// wg.Add(2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			go func (data string)  {
				fmt.Println("Data dari Channel 1", data)
			}(data)
			counter++
		case data := <-channel2:
			go func (data string)  {
				fmt.Println("Data dari Channel 2", data)
			}(data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break;
		}
	}

	time.Sleep(3 * time.Second)

}