package program3

import "fmt"

// send only channerl passed in goroutine
func sendData(sendCh chan<- int) {
	sendCh <- 10
}

func receiveData(receiveCh <-chan int) {
	fmt.Println(<-receiveCh)
}

func producer(chnl chan<- int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func receiver(chnl <-chan int) {
	for {
		v, ok := <-chnl
		if !ok {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

func UniAndBiChannel() {
	// sendCh := make(<-chan int) // receive only channel (unidirectional)
	// sendCh := make(chan<- int) // send only channel (unidirectional)
	sendCh := make(chan int) // bidirection channel
	loopCh := make(chan int)

	go sendData(sendCh)
	go producer(loopCh)

	receiveData(sendCh)
	receiver(loopCh)

}
