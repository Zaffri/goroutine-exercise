package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func startSelect() {
	fmt.Println("SELECT STARTING...")
	myChannel := make(chan string)

	go selectFunc(myChannel)

	select {
	case data := <-myChannel:
		fmt.Printf("Data recieved from channel: %s\n", data)

	case <-time.After(time.Microsecond * 2000):
		fmt.Println("Timeout")
	}

	fmt.Println("SELECT FINISHED.")
}

func selectFunc(myChannel chan string) {
	seconds := rand.IntN(3)
	time.Sleep(time.Duration(seconds) * time.Second)
	myChannel <- "HELLO"
}
