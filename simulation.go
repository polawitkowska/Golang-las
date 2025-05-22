package main

import "math/rand"

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
							grid[newI][newJ].isBurned = true
						}
					}
				}
			}
		}
	}
}
