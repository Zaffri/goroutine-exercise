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

	case <-time.After(time.Second * 2):
		fmt.Println("Timeout")
	}

	fmt.Printf("SELECT FINISHED.\n\n")
}

func selectFunc(myChannel chan string) {
	seconds := rand.IntN(3)
	fmt.Printf("Sleeping for %d seconds\n", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	myChannel <- "HELLO"
}
