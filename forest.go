package main

import (
	"fmt"
	"math/rand"
)

type forest struct {
	isTree   bool
	isBurned bool
}

func makeForest(x int, y int, p float64) [][]forest {
	grid := make([][]forest, x)
	for i := range grid {
		grid[i] = make([]forest, y)
	}

	howManyTrees := (float64(x) * float64(y)) * (p / 100)

	count := 0
	for count < int(howManyTrees) {
		i := rand.Intn(x)
		j := rand.Intn(y)

		if !grid[i][j].isTree {
			grid[i][j].isTree = true
			count++
		}
	}

	return grid
}

func countTrees(grid [][]forest) int {
	howManyTrees := 0

	for _, row := range grid {
		for _, cell := range row {
			if cell.isTree && !cell.isBurned {
				howManyTrees++
			}
		}
	}

	return howManyTrees
}

func printForest(grid [][]forest) {
	for _, row := range grid {
		fmt.Println(" ")
		for _, cell := range row {
			if cell.isTree && !cell.isBurned {
				fmt.Print("T ")
			} else if cell.isTree && cell.isBurned {
				fmt.Print("B ")
			} else {
				fmt.Print("_ ")
			}
		}
	}
	fmt.Println(" ")
}
