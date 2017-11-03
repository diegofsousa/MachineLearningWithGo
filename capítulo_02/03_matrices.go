package main

import (
	"fmt"
	//"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	components := []float64{1.2, -5.7, -2.4, 7.3}
	a := mat.NewDense(2, 2, components)
	fa := mat.Formatted(a, mat.Prefix(" "))

	fmt.Printf(" %v\n\n", fa)

}
