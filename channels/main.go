package main

import (
	"fmt"
	"time"
)

// Sum result by using Chananl
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult

}

// Go Routine sync
func processing(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Processing...")
}

// Sanding Email by using Chananl
func emailSender(emailCha <-chan string, done chan<- bool) {
	for email := range emailCha {
		fmt.Println("Sending email to...", email)
		time.Sleep(time.Second)
	}
	defer func() { done <- true }()

}

func main() {
	result := make(chan int)
	go sum(result, 3, 5)
	res := <-result
	fmt.Println(res)

	don := make(chan bool)
	go processing(don)

	// Sending email
	emailChan := make(chan string, 5)
	// emailChan <- "ashish@gmail.com"
	// emailChan <- "ashish.dev@gmail.com"

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	emailSender(emailChan, don)
	close(emailChan)
	<-don

	////////////////////////////////////////
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "ashish"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Value := <-chan1:
			fmt.Println("Chan 1 Value recive.. ", chan1Value)

		case chan2Value := <-chan2:
			fmt.Println("Chan 2 Value recive.. ", chan2Value)
		}
	}

}
