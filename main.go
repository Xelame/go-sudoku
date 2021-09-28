package main

import (
	"fmt"
	"os"
)

func main() {
	tableau := [][]rune{}
	for i := range os.Args[1:] {
		tableau = append(tableau, []rune(os.Args[1:][i]))
	}
	IsValid(tableau, 0)
	for _, value := range tableau {
		fmt.Println(string(value))
	}
}

func ValidGrid(tableau [][]rune) bool {
	iscorect := true
	if len(tableau) == 9 {
		for _, ligne := range tableau {
			for _, colonne := range ligne {
				if colonne != 46 {
					if colonne < 49 || colonne > 58 {
						iscorect = false
					}
				}
			}
		}
	} else {
		iscorect = false
	}
	return iscorect
}

func LineMissing(board [][]rune, k, line int) bool {
	for column := 0; column < 9; column++ {
		if board[line][column] == rune(k)+48 {
			return false
		}
	}
	return true
}

func ColumnMissing(board [][]rune, k, column int) bool {
	for line := 0; line < 9; line++ {
		if board[line][column] == rune(k)+48 {
			return false
		}
	}
	return true
}

func BlocMissing(board [][]rune, k, j, i int) bool {
	i2 := i - (i % 3)
	j2 := j - (j % 3)
	for i = i2; i < i2+3; i++ {
		for j = j2; j < j2+3; j++ {
			if board[i][j] == rune(k)+48 {
				return false
			}
		}
	}
	return true
}

func IsValid(board [][]rune, position int) bool {

	// Si on est à la 82e case (on sort du tableau)
	if position == 9*9 {
		return true
	}
	// On récupère les coordonnées de la case
	i := position % 9
	j := position / 9

	// Si la case n'est pas un point, on passe à la suivante (appel récursif)
	if board[i][j] != 46 {
		return IsValid(board, position+1)
	}
	// énumération des valeurs possibles
	for k := 1; k <= 9; k++ {
		// Si la valeur est absente, donc autorisée
		if LineMissing(board, k, i) && ColumnMissing(board, k, j) && BlocMissing(board, k, i, j) {
			// On enregistre k dans la grille
			fmt.Println(k)
			board[i][j] = rune(k) + 48
			// On appelle récursivement la fonction estValide(), pour voir si ce choix est bon par la suite
			if IsValid(board, position+1) {
				return true // Si le choix est bon, plus la peine de continuer, on renvoie true :)
			}
		}
	}
	// Tous les chiffres ont été testés, aucun n'est bon, on réinitialise la case
	board[i][j] = 46
	// Puis on retourne false :(
	return false
}

// Aide pour afficher notre sudoku
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
