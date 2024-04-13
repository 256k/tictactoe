package main

import (
	"fmt"
	// "os"
)

var grid = [3][3]int{
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
}

var gridchar = []string{"-", "X", "O"}
var numberOfPlayers = 2

func main() {
	fmt.Printf("Tic Tac Toe\n============\n")
	fmt.Printf("enter the coordinates to fill, ex: 2 1\n\n")
	drawGrid()
	gameLoop()
}

func gameLoop() {
	var x, y int
	player := 1
	for {
		fmt.Printf("Player %d move: ", player)
		fmt.Scanf("%d %d\n", &x, &y)
		if x > 2 || y > 2 {
			fmt.Printf("The coordinates are out of range\nPlease try again\n")
			continue
		}
		if grid[x][y] != 0 {
			fmt.Printf("That slot is already filled\nPlease try again\n")
			continue
		}
		grid[x][y] = player
		player++
		if player == numberOfPlayers+1 {
			player = 1
		}
		drawGrid()
	}
}

func drawGrid() {
	fmt.Printf("\n=============\n")
	for i := 0; i < len(grid); i++ {
		fmt.Printf("| ")
		for j := 0; j < len(grid); j++ {
			fmt.Printf("%s", gridchar[grid[j][i]])
			if j < len(grid[i])-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Printf(" |\n")
		if i < len(grid)-1 {
			fmt.Printf("-------------\n")
		}
	}
	fmt.Printf("=============\n")
}
