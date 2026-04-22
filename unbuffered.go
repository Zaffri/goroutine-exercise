package main

import (
	"fmt"
	"time"
)

func startUnbuffered() {
	fmt.Println("UNBUFFERED STARTING...")
	myChannel := make(chan string)

	go doSomething(myChannel)

	myChannel <- "SEND_123"
	messageRecieved := <-myChannel

	fmt.Printf("Messaged recieved from unbuffered channel: %s\n", messageRecieved)
	fmt.Printf("UNBUFFERED FINISHED.\n\n")
}

func doSomething(myChannel chan string) {
	fmt.Println("doSomething is waiting for myChannel message")
	recieved := <-myChannel

	fmt.Printf("doSomething has recieved myChannel message: %s\n", recieved)
	time.Sleep(time.Second)

	myChannel <- "RETURN_123"
}
