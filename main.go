package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
)

func main() {
	x1 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y1 := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	x2 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y2 := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}

	x3 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y3 := []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}

	x4 := []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
	y4 := []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}

	// Perform linear regression on x1 and y1
	slope1, intercept1, err := linearRegression(x1, y1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Linear Regression for x1 and y1: y = %fx + %f\n", slope1, intercept1)

	// Perform linear regression on x2 and y2
	slope2, intercept2, err := linearRegression(x2, y2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Linear Regression for x2 and y2: y = %fx + %f\n", slope2, intercept2)

	// Perform linear regression on x3 and y3
	slope3, intercept3, err := linearRegression(x3, y3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Linear Regression for x3 and y3: y = %fx + %f\n", slope3, intercept3)

	// Perform linear regression on x4 and y4
	slope4, intercept4, err := linearRegression(x4, y4)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Linear Regression for x4 and y4: y = %fx + %f\n", slope4, intercept4)
}

func linearRegression(x, y []float64) (float64, float64, error) {
	if len(x) != len(y) {
		return 0, 0, fmt.Errorf("input slices must have the same length")
	}

	xMean, err := stats.Mean(x)
	if err != nil {
		return 0, 0, err
	}
	yMean, err := stats.Mean(y)
	if err != nil {
		return 0, 0, err
	}

	var numerator, denominator float64
	for i := 0; i < len(x); i++ {
		numerator += (x[i] - xMean) * (y[i] - yMean)
		denominator += (x[i] - xMean) * (x[i] - xMean)
	}

	if denominator == 0 {
		return 0, 0, fmt.Errorf("denominator is zero, cannot compute slope")
	}

	slope := numerator / denominator
	intercept := yMean - slope*xMean

	return slope, intercept, nil
}
