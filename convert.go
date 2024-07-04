package main

import "fmt"

func convertCelsius(kelvin float64) float64 {

	return kelvin - 273

}

func main() {

	test := convertCelsius(372.2)
	fmt.Printf("%.2f", test)

}
