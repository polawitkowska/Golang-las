package main

import (
	"fmt"
	"math/rand"
)

//Typ forest
type forest struct {
	isTree   bool
	isBurned bool
	humidity int //Dodałam wilgotność dla każdego drzewa
}

//Tworzenie lasu
func makeForest(x int, y int, p float64) [][]forest {
	//grid to nasz las
	grid := make([][]forest, x)
	for i := range grid {
		grid[i] = make([]forest, y)
	}

	howManyTrees := (float64(x) * float64(y)) * (p / 100) // Program liczy ile drzew utworzyć dla naszego rozmiaru lasu, aby zajmowały p%

	count := 0	// Licznik ile drzew zostało już utworzonych
	// Zamiast przechodzić po każdej komórce grid i losować czy tam będzie drzewo, program losuje komórkę
	for count < int(howManyTrees) { // Robi to tak długo aż count będzie równe wcześniej obliczonemu howManyTrees
		i := rand.Intn(x)
		j := rand.Intn(y)

		if !grid[i][j].isTree {
			grid[i][j].isTree = true // Jeżeli wylosowana komórka nie jest jeszcze drzewem, ustawiamy ją na drzewo
			grid[i][j].humidity = rand.Intn(101) // Każdemu drzewu dajemy losowe humidity
			count++
		}
	}

	return grid
}

// Funkcja do liczenia ilości drzew w lesie, przydaje się do obliczenia ile drzew się spaliło itp.
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

// Funkcja która drukuje las w ładny sposób np:
// T _ _ B     T - tree
// _ T _ B     B - burned
// _ _ _ _     _ - puste
// niezbyt dobrze pokazuje las przy dużych x i y
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
