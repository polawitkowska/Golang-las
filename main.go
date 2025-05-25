package main

import (
	"flag"
	"fmt"
)

func getParameters() (int, int, float64) {
	var x = flag.Int("x", 100, "Width of the grid")
    var y = flag.Int("y", 100, "Height of the grid")
    var p = flag.Float64("p", 40.0, "Density of trees")

    flag.Parse()

    if *p > 100 {
        *p = 100
    }

    if *x <= 0 || *y <= 0 || *p <= 0 {
        fmt.Println("Parameters must be positive numbers!")
        return 0, 0, 0
    }

    fmt.Printf("Running simulation with parameters: x=%d, y=%d, p=%.2f\n", *x, *y, *p)
    return *x, *y, *p
}

func main() {
	x, y, p := getParameters()

	if x == 0 && y == 0 && p == 0 {
		return
	}

	//grid := makeForest(x, y, p)
	//fmt.Println("Forest before thunder: ")
	//printForest(grid)
	//
	//thunder(x, y, grid)
	//fmt.Println("Forest after thunder: ")
	//printForest(grid)

	result := simulation(x, y, p)
	fmt.Printf("%d,%d,%.f,%.2f", x, y, p, result)
}
