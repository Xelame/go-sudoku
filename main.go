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


func SplitWhiteSpaces(s string) []string {
	tab := []string{}
	cont := -1
	if s[0] == ' ' {
		cont = 0
	}
	for i := 1; i < len(s); i++ {
		if s[i] == ' ' {
			if s[cont+1:i] != "" {
				tab = append(tab, s[cont+1:i])
			}
			cont = i
		}
	}
	tab = append(tab, s[cont+1:])
	return tab
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