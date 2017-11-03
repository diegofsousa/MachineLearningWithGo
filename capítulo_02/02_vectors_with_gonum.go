package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	myvector := mat.NewVecDense(2, []float64{11.0, 5.2})
	fmt.Println(myvector)

	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}

	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("O dot product de A e B é: %0.2f\n", dotProduct)

	floats.Scale(1.5, vectorA)
	fmt.Printf("Escalando A com 1.5: %v\n", vectorA)

	normB := floats.Norm(vectorB, 2)
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)
}
