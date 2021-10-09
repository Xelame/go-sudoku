### RAID SUDOKU

Hi, This program is a raid during a school parkour.

The goal of this raid is to resolve a sudoku given as arguments in 2 days with a particular structure in go.

For exemple if I write in my bash console (linux/WSL) :

go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"  
(To explain, I compile and run my main.go file with 9 arguments who I give which is a string of 9 digits or a dot to stand for an empty case)

The structure of sudoku therefore looks like this :

`╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗`  
`║ . │ 9 │ 6 ║ . │ 4 │ . ║ . │ . │ 1 ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ 1 │ . │ . ║ . │ 6 │ . ║ . │ . │ 4 ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ 5 │ . │ 4 ║ 8 │ 1 │ . ║ 3 │ 9 │ . ║`  
`╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣`  
`║ . │ . │ 7 ║ 9 │ 5 │ . ║ . │ 4 │ 3 ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ . │ 3 │ . ║ . │ 8 │ . ║ . │ . │ . ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ 4 │ . │ 5 ║ . │ 2 │ 3 ║ . │ 1 │ 8 ║`  
`╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣`  
`║ . │ 1 │ . ║ 6 │ 3 │ . ║ . │ 5 │ 9 ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ . │ 5 │ 9 ║ . │ 7 │ . ║ 8 │ 3 │ . ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ . │ . │ 3 ║ 5 │ 9 │ . ║ . │ . │ 7 ║`  
`╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝`  

and a exemple of solution can be this table :  

`╔═══╤═══╤═══╦═══╤═══╤═══╦═══╤═══╤═══╗`  
`║ 3 │ 9 │ 6 ║ 2 │ 4 │ 5 ║ 7 │ 8 │ 1 ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ 1 │ 7 │ 8 ║ 3 │ 6 │ 9 ║ 5 │ 2 │ 4 ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ 5 │ 2 │ 4 ║ 8 │ 1 │ 7 ║ 3 │ 9 │ 6 ║`  
`╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣`  
`║ 2 │ 8 │ 7 ║ 9 │ 5 │ 1 ║ 6 │ 4 │ 3 ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ 9 │ 3 │ 1 ║ 4 │ 8 │ 6 ║ 2 │ 7 │ 5 ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ 4 │ 6 │ 5 ║ 7 │ 2 │ 3 ║ 9 │ 1 │ 8 ║`  
`╠═══╪═══╪═══╬═══╪═══╪═══╬═══╪═══╪═══╣`  
`║ 7 │ 1 │ 2 ║ 6 │ 3 │ 8 ║ 4 │ 5 │ 9 ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ 6 │ 5 │ 9 ║ 1 │ 7 │ 4 ║ 8 │ 3 │ 2 ║`  
`╟───┼───┼───╫───┼───┼───╫───┼───┼───╢`  
`║ 8 │ 4 │ 3 ║ 5 │ 9 │ 2 ║ 1 │ 6 │ 7 ║`  
`╚═══╧═══╧═══╩═══╧═══╧═══╩═══╧═══╧═══╝`  