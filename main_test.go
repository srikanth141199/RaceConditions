package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	os.Stdout = w

	outputCh := make(chan string)

	go func() {
		result, err := io.ReadAll(r)
		if err != nil {
			t.Error(err)
			return
		}

		outputCh <- string(result)
	}()

	main()

	w.Close()
	os.Stdout = stdOut

	output := <-outputCh

	if !strings.Contains(output, "$34320.00") {
		t.Errorf(
			"Expected output to contain 'Final bank balance: $34320.00', but got: %s",
			output,
		)
	}
}
