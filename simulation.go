package main

import (
	"fmt"
	"math/rand"
)

func thunder(x int, y int, grid [][]forest) {
	targetX := rand.Intn(x)
	targetY := rand.Intn(y)

	if !grid[targetX][targetY].isTree {
		return
	}

	grid[targetX][targetY].isBurned = true

	for _ = range x {
		for i := range x {
			for j := range y {
				if grid[i][j].isTree && grid[i][j].isBurned {
					directions := [][2]int{
						{-1, 0}, {1, 0}, {0, -1}, {0, 1},
						{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
					}

					for _, direction := range directions {
						newI, newJ := i+direction[0], j+direction[1]
						if newI >= 0 && newI < x && newJ >= 0 && newJ < y &&
							grid[newI][newJ].isTree && !grid[newI][newJ].isBurned {

							humidity := grid[newI][newJ].humidity
							chance := rand.Intn(101)

							if chance > humidity {
								grid[newI][newJ].isBurned = true
							}
						}
					}
				}
			}
		}
	}
}

func simulation(x int, y int, p float64) float64 {
	var n = 50

	treesCount := make([]int, n)
	for i := range n {
		grid := makeForest(x, y, p)
		howManyTreesBefore := countTrees(grid)
		thunder(x, y, grid)
		howManyTreesAfter := countTrees(grid)

		treesCount[i] = howManyTreesBefore - howManyTreesAfter
	}

	sum := 0
	for _, v := range treesCount {
		sum += v
	}
	average := sum / n
	averagePercent := (float64(average) / float64(x*y)) * 100

	fmt.Printf("For density of forest %.2f%% on average trees burnt %.2f%%\n", p, averagePercent)
	return averagePercent
}
