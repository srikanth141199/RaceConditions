package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello World!"
	wg.Add(2)
	go updateMessage("GoodBye 1!")
	go updateMessage("GoodBye 2!")
	wg.Wait()

	if msg != "GoodBye 1!" {
		t.Errorf("Expected msg to be 'GoodBye 1!', but got '%s'", msg)
	}
}
