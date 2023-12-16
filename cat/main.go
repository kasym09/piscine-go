package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	size := len(os.Args)

	for i := 1; i < size; i++ {
		data, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			io.WriteString(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		}

		io.WriteString(os.Stdout, string(data))
	}
}
