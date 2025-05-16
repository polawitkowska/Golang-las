package main

import (
	"fmt"
	"math/rand"
)

type forest struct {
	isTree   bool
	isBurned bool
}

func getParameters() [3]int {
	var parameters [3]int
	fmt.Print("Enter X, Y and P:\n")
	_, err := fmt.Scanln(&parameters[0], &parameters[1], &parameters[2])
	if err != nil {
		fmt.Println("Error while getting parameters. Try again.")
		return [3]int{0, 0, 0}
	}

	if parameters[0] <= 0 || parameters[1] <= 0 || parameters[2] <= 0 {
		fmt.Println("Parameters X, Y and P cannot be 0 or negative! Try again.")
		return [3]int{0, 0, 0}
	}

	//if parameters[2] > 100 {
	//	fmt.Println("P must be less or equal to 100!")
	//	return [3]int{0, 0, 0}
	//}

	return parameters
}

func makeForest(x int, y int, p int) [][]forest {
	grid := make([][]forest, x)
	for i := range grid {
		grid[i] = make([]forest, y)
	}

	howManyTrees := (float64(x) * float64(y)) * (float64(p) / 100)

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

func main() {
	parameters := getParameters()

	if parameters[0] == 0 && parameters[1] == 0 && parameters[2] == 0 {
		return
	}

	x := parameters[0]
	y := parameters[1]
	p := parameters[2]

	var n = 100

	treesCount := make([]int, n)
	for i := range n {
		grid := makeForest(x, y, p)
		thunder(x, y, grid)
		howManyTrees := countTrees(grid)

		treesCount[i] = howManyTrees
	}

	sum := 0
	for _, v := range treesCount {
		sum += v
	}
	average := sum / n
	averagePercent := (float64(average) / float64(x*y)) * 100

	fmt.Println("For density of forest", p, "% on average trees left", int(averagePercent), "%")
}
