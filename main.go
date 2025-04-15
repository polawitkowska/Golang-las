package main

import (
	"fmt"
	"math/rand"
)

type forest struct {
	isTree    bool
	isBurning bool
	isBurned  bool
}

func main() {
	grid := make([][]forest, 6)
	for i := range grid {
		grid[i] = make([]forest, 6)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = forest{
				isTree:    rand.Intn(2) == 1,
				isBurning: false,
				isBurned:  false,
			}
		}
	}

	for i, row := range grid {
		for j, cell := range row {
			fmt.Println("Element at", i, j, "isTree:", cell.isTree)
		}
	}
}
