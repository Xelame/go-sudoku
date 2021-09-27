package main

import (
	"os"

	"fmt"
)

func main() {
	tableau := os.Args[1:]
	for _, ligne := range tableau {
		colonnes := [9]string{}
		for index, colonne := range ligne {
			colonnes[index] = string(colonne)
		}
		ligne := Join(colonnes, " ")
		fmt.Println(ligne)
	}
}

/*
func TestNumber() {
	arguments := os.Args[1:]
	for _, value := range arguments {
		if value == "." {
			betting := true
			for betNumber := '1'; betNumber < '9' && betting; betNumber++ {
				value
			}
		}

	}
}
*/
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
