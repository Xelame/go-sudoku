package main

import (
	"os"

	"fmt"
)

func main() {
	tableau := os.Args[1:]
	if len(tableau) == 9 {
		for _, ligne := range tableau {
			colonnes := [9]string{}
			for index, colonne := range ligne {
				colonnes[index] = string(colonne)
			}
			ligne := Join(colonnes, " ")
			fmt.Println(ligne)
		}
	} else {
		fmt.Println("ERROR")
	}
}

func TestNumber(betNumber string) {
	tableau := os.Args[1:]
	for _, ligne := range tableau {
		for index, value := range ligne {
			if value == '.' {
				betting := true
				for betNumber := '1'; betNumber < '9' && betting; betNumber++ {
					value = betNumber
					for indextest, testValue := range ligne {

					}
				}
			}
		}
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
