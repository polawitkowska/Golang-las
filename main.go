package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type forest struct {
	isTree   bool
	isBurned bool
}

func getParameters() (int, int, float64) {
	x := 20
	y := 20
	p := 25.0
	var input string

	fmt.Printf("Enter X (currently %d) or press Enter to skip: ", x)
	if _, err := fmt.Scanln(&input); err == nil {
		if newX, err := strconv.Atoi(input); err == nil {
			if newX <= 0 {
				fmt.Println("X cannot be 0 or negative! Using default value.")
			} else {
				x = newX
			}
		}
	}

	fmt.Printf("Enter Y (currently %d) or press Enter to skip: ", y)
	if _, err := fmt.Scanln(&input); err == nil {
		if newY, err := strconv.Atoi(input); err == nil {
			if newY <= 0 {
				fmt.Println("Y cannot be 0 or negative! Using default value.")
			} else {
				y = newY
			}
		}
	}

	fmt.Printf("Enter P (currently %.2f) or press Enter to skip: ", p)
	if _, err := fmt.Scanln(&input); err == nil {
		if newP, err := strconv.ParseFloat(input, 64); err == nil {
			if newP <= 0 {
				fmt.Println("P cannot be 0 or negative! Using default value.")
			} else {
				p = newP
			}
		}
	}

	fmt.Println("x, y and p are:", x, y, p)
	return x, y, p
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
	x, y, p := getParameters()

	if x == 0 && y == 0 && p == 0 {
		return
	}

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

	fmt.Println("For density of forest", p, "% on average trees burnt", int(averagePercent), "%")
}
