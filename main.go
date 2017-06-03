package main

import (
	"bytes"
	"io/ioutil"
	"os"
)

func main() {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	splitted := bytes.Split(stdin, []byte("\n"))
	if len(splitted) > 1 && string(splitted[len(splitted)-1]) == "" {
		splitted = splitted[:len(splitted)-2]
	}
	os.Stdout.Write([]byte("["))
	for index, line := range splitted {
		if index == len(splitted)-1 {
			os.Stdout.Write(line)
			os.Stdout.Write([]byte("]"))
		} else {
			os.Stdout.Write(line)
			os.Stdout.Write([]byte(","))
		}
	}
	os.Stdout.Write([]byte("\n"))
}
