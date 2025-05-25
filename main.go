package main

import (
	"flag"
	"fmt"
)

// Pobieranie parametrów używając flag - łatwiej jak chce się odpalić program pare razy z tymi samymi parametrami
func getParameters() (int, int, float64) {
	var x = flag.Int("x", 100, "Width of the grid")
    var y = flag.Int("y", 100, "Height of the grid")
    var p = flag.Float64("p", 40.0, "Density of trees") // Domyślne parametry to 100 100 40%

    flag.Parse()

    if *p > 100 {
        *p = 100
    }

    if *x <= 0 || *y <= 0 || *p <= 0 {
        fmt.Println("Parameters must be positive numbers!")
        return 0, 0, 0
    }

    fmt.Printf("Parameters are: x=%d, y=%d, p=%.2f\n", *x, *y, *p)
    return *x, *y, *p
}

func main() {
	x, y, p := getParameters()

	if x == 0 && y == 0 && p == 0 {
		return
	}

	//Jedno uruchomienie spalania lasu (funkcja thunder), terminal pokazuje jak las wygląda przed i po uderzeniu pioruna
	//grid := makeForest(x, y, p)
	//fmt.Println("Forest before thunder: ")
	//printForest(grid)
	//
	//thunder(x, y, grid)
	//fmt.Println("Forest after thunder: ")
	//printForest(grid)

	//Uruchomienie funkcji "simulation" która uruchamia spalanie lasu 50 razy i zwraca średnią ilość spalonych drzew w %
	result := simulation(x, y, p)
	fmt.Printf("%d,%d,%.f,%.2f", x, y, p, result)
}
