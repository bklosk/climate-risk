package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMain_PrintHelloWorld(t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old // restoring the real stdout

	out := <-outC

	expected := "heyo, world.\n"
	if out != expected {
		t.Errorf("Expected output %q, but got %q", expected, out)
	}
}
