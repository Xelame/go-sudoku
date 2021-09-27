package main

import (
	"os"

	"fmt"
)

func main() {
	arguments := os.Args[1:]
	for _, arg := range arguments {
		fmt.Println(arg)
	}
}
