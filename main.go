package main

import (
	"fmt"
	"os"
)

func main() {
	board := [][]rune{}
	for i := range os.Args[1:] {
		board = append(board, []rune(os.Args[1:][i]))
	}
	fmt.Println("Sudoku donnée : ")
	Display(board)
	fmt.Println("Sudoku résolu : ")
	NumberIsValid(board, 0)
	Display(board)
}

// Function to test if a number is egal to my bet in a same line
func LineMissing(board [][]rune, k, line int) bool {
	for column := 0; column < 9; column++ {
		if board[line][column] == rune(k)+48 { // line is const and column increase
			return false
		}
	}
	return true
}

// Function to test if a number is egal to my bet in a same column
func ColumnMissing(board [][]rune, k, column int) bool {
	for line := 0; line < 9; line++ {
		if board[line][column] == rune(k)+48 { // column is const and line increase
			return false
		}
	}
	return true
}

// Function to test if a number is egal to my bet in a same block
func BlocMissing(board [][]rune, k, column, line int) bool {
	lineBlock := line - (line % 3)       // 3 block cut, 0-2 3-5 6-8
	columnBlock := column - (column % 3) // 3 block cut, 0-2 3-5 6-8
	for line = lineBlock; line < lineBlock+3; line++ {
		for column = columnBlock; column < columnBlock+3; column++ {
			if board[line][column] == rune(k)+48 { // Scan the block entire
				return false
			}
		}
	}
	return true
}

// Recursive function which travel my sudoku
func NumberIsValid(board [][]rune, position int) bool {

	// Stop on the 82th case (out of range)
	if position == 9*9 {
		return true
	}
	// We take case's coordonates
	line := position / 9
	column := position % 9

	// If is ont a dot, we move to the next one
	if board[line][column] != 46 {
		return NumberIsValid(board, position+1)
	}
	// enumeration of possible values
	for k := 1; k <= 9; k++ {
		// If the value is missing, we can choose them
		if LineMissing(board, k, line) && ColumnMissing(board, k, column) && BlocMissing(board, k, column, line) {
			// We save k in the grid
			board[line][column] = rune(k) + 48
			// We recursively call the function NumberIsValid(), to see if this choice is correct afterwards
			if NumberIsValid(board, position+1) {
				return true // If the choice is good, no more bother to continue, we return true 😁
			}
		}
		// All the figures have been tested, none is good, we reset the box 😔
		board[line][column] = 46
	}
	// And we return false :(
	return false
}

// Help us to print our board
func Join(strs string, sep string) string {
	var str string
	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 {
			str += string(strs[i]) + sep
		} else {
			str += string(strs[i])
		}
	}
	return str
}

// Display our beautiful sudoku (Thanks Maxime 😁)
func Display(board [][]rune) {
	fmt.Println("╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗")
	for i, value := range board {
		fmt.Println("║ " + Join(string(value)[:3], " │ ") + " ║ " + Join(string(value)[3:6], " │ ") + " ║ " + Join(string(value)[6:9], " │ ") + " ║")
		if i < 8 {
			if i == 2 || i == 5 {
				fmt.Println("╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣")
			} else {
				fmt.Println("╟───┼───┼───╫───┼───┼───╫───┼───┼───╢")
			}
		}
	}
	fmt.Println("╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝")
}
