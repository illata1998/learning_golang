package main

import "fmt"

func main() {
	const USDToEUR = 0.8581
	const USDToRUB = 77.96
	const EURToRUB = USDToRUB / USDToEUR
	fmt.Print(EURToRUB)
}