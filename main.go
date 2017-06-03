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
	output := []byte("[")
	for index, line := range splitted {
		if index == len(splitted)-1 {
			output = append(output, line...)
			output = append(output, []byte("]")...)
		} else {
			output = append(output, line...)
			output = append(output, []byte(",")...)
		}
	}
	output = append(output, []byte("\n")...)
	os.Stdout.Write(output)
}
