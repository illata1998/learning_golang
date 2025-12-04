package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPower float64 = 2
	var userHeight = 1.8
	var userWeight float64 = 100
	var BMI = userWeight / math.Pow(userHeight, BMIPower)
	fmt.Print(BMI)
}
