package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var msg string

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Hello World!"

	wg.Add(2)
	go updateMessage("Hello Universe!")
	go updateMessage("Hello Cosmos!")
	wg.Wait()
	fmt.Println(msg)
}
