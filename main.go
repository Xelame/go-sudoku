package main

import (
	"os"

	"fmt"
)

func main() {
	tableau := os.Args[1:]
	iscorect := true
	if len(tableau) == 9 {
		for _, ligne := range tableau {
			for _, colonne := range ligne {
				if colonne != '.' {
					if  colonne < 48 || colonne > 57 {
						iscorect = false
					}
				}
			}
		}
		if iscorect {
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
	} else {
		fmt.Println("ERROR")
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
