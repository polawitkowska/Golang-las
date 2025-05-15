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

	for a := range int(howManyTrees) {
		i := rand.Intn(x)
		j := rand.Intn(y)
		fmt.Println(a)

		if !grid[i][j].isTree {
			grid[i][j].isTree = true
		} else {
			howManyTrees++
		}

	}

	//for i := range grid {
	//	for j := range grid[i] {
	//		grid[i][j] = forest{
	//			isTree:   rand.Intn(2) == 1,
	//			isBurned: false,
	//		}
	//	}
	//}

	return grid
}

func thunder(x int, y int, grid [][]forest) {
	targetX := rand.Intn(x)
	targetY := rand.Intn(y)

	fmt.Println("\nTarget:", targetX, targetY)

	if !grid[targetX][targetY].isTree {
		fmt.Println("Target was empty. Repeating.")
		thunder(x, y, grid)
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
		printForest(grid)
	}
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

	grid := makeForest(x, y, p)
	fmt.Print("Forest: ")
	printForest(grid)

	//thunder(x, y, grid)
	//fmt.Println("Forest after thunder: ")
	//printForest(grid)
}
