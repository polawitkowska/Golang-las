package main

import (
	"fmt"
	"math/rand"
)

// Funkcja "pioruna", który uderza w losowe pole w lesie
func thunder(x int, y int, grid [][]forest) {
	targetX := rand.Intn(x)
	targetY := rand.Intn(y)

	if !grid[targetX][targetY].isTree {
		return
	}

	grid[targetX][targetY].isBurned = true

	// Logika palenia drzew obok wybranego na początku drzewa
	// To była najtrudniejsza część zadania
	for _ = range x {  // Logika palenia powtarza się x razy, jak powtarzała się za mało razy to część lasu się nie paliła, a powinna
		for i := range x {
			for j := range y {
				if grid[i][j].isTree && grid[i][j].isBurned {
					directions := [][2]int{ 		// Zamiast pisać dużo if'ów sprawdzających, czy grid[i][j-1], grid[i-1][j] etc istnieje i czy jest niespalonym drzewem
						{-1, 0}, {1, 0}, {0, -1}, {0, 1},	// mam tablicę z przesunięciami na boki i na ukos
						{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
					}

					for _, direction := range directions {						// Dla każdego elementu tablicy z przesunięciami
						newI, newJ := i+direction[0], j+direction[1]			// tworzę "nowe" przesunięte indeksy
						if newI >= 0 && newI < x && newJ >= 0 && newJ < y &&	// i sprawdzam, czy te indeksy istnieją i są niespalonym drzewem
							grid[newI][newJ].isTree && !grid[newI][newJ].isBurned {
																	// Jeżeli tak, to bez humidity drzewo byłoby po prostu spalone
							humidity := grid[newI][newJ].humidity	// ale z humidity mam 100-humidity% szans na to, że drzewo się spali
							chance := rand.Intn(101)				// czyli losuję liczbę od 0 do 100
																	
							if chance > humidity {					// i jeżeli wylosowana liczba jest większa od humidity drzewa
								grid[newI][newJ].isBurned = true	// to drzewo się pali
							}
						}
					}
				}
			}
		}
	}
}

// Funkcja simulation
func simulation(x int, y int, p float64) float64 {
	var n = 50

	treesCount := make([]int, n)
	for i := range n { // Funckja puszcza funkcję thunder dla losowo stworzonego lasu o parametrach x y p 50 razy
		grid := makeForest(x, y, p)
		howManyTreesBefore := countTrees(grid)
		thunder(x, y, grid)
		howManyTreesAfter := countTrees(grid)

		treesCount[i] = howManyTreesBefore - howManyTreesAfter // zapisuje ile drzew się spaliło i zwraca ten wynik
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
