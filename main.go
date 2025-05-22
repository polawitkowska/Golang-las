package main

import (
	"fmt"
	"strconv"
)

func getParameters() (int, int, float64) {
	x := 100
	y := 100
	p := 30.0
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

	fmt.Printf("For density of forest %.2f%% on average trees burnt %.2f%%\n", p, averagePercent)
}
