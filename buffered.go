package main

import (
	"fmt"
)

func startBuffered() {
	fmt.Println("BUFFERED STARTING...")
	myChannel := make(chan string, 3)
	go doSomethingElse(myChannel)

	for message := range myChannel {
		fmt.Printf("Received: %s\n", message)
	}

	fmt.Println("BUFFERED FINISHED.")
}

func doSomethingElse(myChannel chan string) {
	for x := 1; x <= 5; x++ {
		message := fmt.Sprintf("Item %d", x)
		myChannel <- message
		fmt.Printf("Sent: %s\n", message)
	}

	close(myChannel)
}
