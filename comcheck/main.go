package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	for _, i := range arguments {
		if i == "01" || i == "galaxy" || i == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
