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

func makeForest(n int) [][]forest {
	grid := make([][]forest, n)
	for i := range grid {
		grid[i] = make([]forest, n)
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

	return grid
}

func thunder(n int, grid [][]forest) {
	targetX := rand.Intn(n)
	targetY := rand.Intn(n)

	fmt.Println("\nTarget:", targetX, targetY)

	if !grid[targetX][targetY].isTree {
		fmt.Println("Target was empty")
		return
	}

	grid[targetX][targetY].isBurning = true

	for _ = range n {
		for i := range n {
			for j := range n {
				if grid[i][j].isTree && grid[i][j].isBurning && !grid[i][j].isBurned {
					directions := [][2]int{
						{-1, 0}, {1, 0}, {0, -1}, {0, 1},
						{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
					}

					for _, direction := range directions {
						newI, newJ := i+direction[0], j+direction[1]
						if newI >= 0 && newI < n && newJ >= 0 && newJ < n &&
							grid[newI][newJ].isTree && !grid[newI][newJ].isBurning {
							grid[newI][newJ].isBurning = true
						}
					}
					grid[i][j].isBurned = true
				}
			}
		}
	}
}

func main() {
	n := 6
	grid := makeForest(n)

	fmt.Println("Forest: ")
	for i := range grid {
		fmt.Println(grid[i])
	}

	thunder(n, grid)
	fmt.Println("Forest after thunder: ")
	for i := range grid {
		fmt.Println(grid[i])
	}
}
