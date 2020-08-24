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
	fmt.Print("Number of generations: ")
	fmt.Scanln(&generation)
	for p := 0; p < generation; p++ {
		fmt.Print("Number of rows: ")
		fmt.Scanln(&rows)
		fmt.Print("Number of colums: ")
		fmt.Scanln(&cols)
		fmt.Println(rows, cols, generation)
		//var board [1000][1000]int
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
