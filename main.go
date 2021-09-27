package main

import (
	"os"

	"fmt"
)

func main() {
	arguments := os.Args[1:]
	for _, argu := range arguments {
		fmt.Println(argu)
	}
}
