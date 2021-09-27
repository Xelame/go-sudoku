package main

import (
	"os"

	"fmt"
)

func main() {
	arguments := os.Args[1:]
	for _, arg := range arguments {
		caracteres := make([]string, 9)
		for index,cara := range arg {
			caracteres[index] = string(cara)
		}
		arg := Join(caracteres, " ")
		fmt.Println(arg)
	}
}

func Join(strs []string, sep string) string {
	var str string
	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 {
			str += strs[i] + sep
		} else {
			str += strs[i]
		}
	}
	return str
}