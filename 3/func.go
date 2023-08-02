package main

import "fmt"

func paintNeeded(height float64, width float64) float64 {
	area := height * width
	return area / 10.0
	// fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {

	var amount, total float64

	amount = paintNeeded(4.2, 3.0)
	total += amount
	fmt.Printf("%.2f liters needed\n", amount)

	amount = paintNeeded(5.2, 3.5)
	total += amount
	fmt.Printf("%.2f liters needed\n", amount)
	fmt.Printf("Total: %.2f liters\n", total)

}
