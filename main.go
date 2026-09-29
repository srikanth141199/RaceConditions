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

// func updateMessage(s string, m *sync.Mutex) {
// 	defer wg.Done()
// 	m.Lock()
// 	defer m.Unlock()
// 	msg = s
// }

// func main() {
// 	msg = "Hello World!"

// 	var mutex sync.Mutex

// 	wg.Add(2)
// 	go updateMessage("Hello Universe!", &mutex)
// 	go updateMessage("Hello Cosmos!", &mutex)
// 	wg.Wait()
// 	fmt.Println(msg)
// }
