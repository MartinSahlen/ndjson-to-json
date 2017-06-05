package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	os.Stdout.Write([]byte("["))
	for {
		bytes, err := stdin.ReadBytes(byte('\n'))
		os.Stdout.Write(bytes)
		if err == io.EOF || string(bytes) == "" {
			break
		}
		os.Stdout.Write([]byte(","))
	}
	os.Stdout.Write([]byte("]"))
	os.Stdout.Write([]byte("\n"))
}
