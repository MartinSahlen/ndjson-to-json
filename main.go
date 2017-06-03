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
		if err == io.EOF {
			break
		}
		if string(bytes) == "" {
			continue
		}
		os.Stdout.Write(bytes)
		b, _ := stdin.Peek(1)
		if len(b) > 0 {
			os.Stdout.Write([]byte(","))
		}
	}

	os.Stdout.Write([]byte("]"))
	os.Stdout.Write([]byte("\n"))
}
