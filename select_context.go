package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func startSelectContext() {
	fmt.Println("SELECT W/ CONTEXT STARTING...")
	myContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	myChannel := make(chan string)
	go selectContextFunc(myContext, myChannel)

	select {
	case <-myChannel:
		fmt.Println("Success")
	case <-myContext.Done():
		fmt.Println("Timed out")
	}
	fmt.Printf("SELECT W/ CONTEXT FINISHED.\n\n")
}

func selectContextFunc(myContext context.Context, myChannel chan string) {
	// handling scenario where context may be cancelled and therefore we need to exit go routine to avoid deadlock/memory leak
	// could also help by making channel buffered...
	seconds := rand.IntN(3)
	fmt.Printf("Sleeping for %d seconds\n", seconds)

	select {
	case <-time.After(time.Duration(seconds) * time.Second):
		myChannel <- "Done"
	case <-myContext.Done():
		return
	}
}

// func selectContextFuncOld(myChannel chan string) {
// 	// problem is that the goroutine could still run after myContext timesout and cancel is called
// 	seconds := rand.IntN(3)
// 	fmt.Printf("Sleeping for %d seconds\n", seconds)
// 	time.Sleep(time.Duration(seconds) * time.Second)

// 	myChannel <- "Done"
// }
