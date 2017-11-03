package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	// Define observed and expected values. Most
	// of the time these will come from your
	// data (website visits, etc.).
	observed := []float64{48, 52}
	expected := []float64{50, 50}
	// Calculate the ChiSquare test statistic.
	chiSquare := stat.ChiSquare(observed, expected)
	fmt.Println(chiSquare)

	chiDist := distuv.ChiSquared{
		K:   2.0,
		Src: nil,
	}
	// Calculate the p-value for our specific test statistic.
	pValue := chiDist.Prob(chiSquare)
	// Output the p-value to standard out.
	fmt.Printf("p-value: %0.4f\n\n", pValue)
}
