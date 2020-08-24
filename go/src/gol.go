package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var generation, rows, cols int
	// Number of times board populated
	fmt.Print("Number of generations: ")
	fmt.Scanln(&generation)
	for p := 0; p < generation; p++ {
		// User input for number of rows
		fmt.Print("Number of rows: ")
		fmt.Scanln(&rows)
		// User input for number of cols
		fmt.Print("Number of colums: ")
		fmt.Scanln(&cols)
		board := make([][]int, rows)
		for i := range board {
			board[i] = make([]int, cols)
		}
		// Create board from random values
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				rand.Seed(time.Now().UnixNano())
				board[i][j] = randomInt(0, 2)
			}
		}
		//printBoard(board)
		//playGame(board)
	}
}
