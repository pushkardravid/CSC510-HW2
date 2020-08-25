package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func countLiveCells(board [][]int, i int, j int) int {
	var checkList = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	var count int = 0
	var row, col int
	for k := range checkList {
		row = i + checkList[k][0]
		col = j + checkList[k][1]
		if row >= 0 && row < len(board) && 0 <= col && col < len(board[0]) && board[row][col]&1 == 1 {
			count++
		}
	}
	return count
}

func playGame(board [][]int, generation int) {
	for p := 0; p < generation; p++ {
		r := len(board)
		c := len(board[0])
		future := make([][]int, r)
		for i := range board {
			future[i] = make([]int, c)
		}
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				count := countLiveCells(board, i, j)
				if board[i][j]&1 == 1 {
					if count == 2 || count == 3 {
						future[i][j] = 1
					} else {
						future[i][j] = 0
					}
				} else {
					if count == 3 {
						future[i][j] = 1
					}
				}
			}
		}
		for m := 0; m < r; m++ {
			for n := 0; n < c; n++ {
				board[m][n] = future[m][n]
			}
		}
		fmt.Println("Generation " + strconv.Itoa(p+1))
		printBoard(board)
	}
}

func printBoard(board [][]int) {
	rows := len(board)
	cols := len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(strconv.Itoa(board[i][j]) + " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var generation, rows, cols int
	// Number of times board populated
	fmt.Print("Number of generations: ")
	fmt.Scanln(&generation)
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
	fmt.Println("Original generated board")
	printBoard(board)
	playGame(board, generation)
}
