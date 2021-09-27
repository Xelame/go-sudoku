package main

import (
	"os"

	"fmt"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 9 {
		for _, arg := range arguments {
			caracteres := [9]string{}
			for index, cara := range arg {
				caracteres[index] = string(cara)
			}
			arg := Join(caracteres, " ")
			fmt.Println(arg)
		}
	} else {
		fmt.Println("ERROR")
	}
}

func Join(strs [9]string, sep string) string {
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
