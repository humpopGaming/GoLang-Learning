package main

import(
	"fmt"
)

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func boilingAndFreezing() (float64, float64) {
	return celsiusToFahrenheit(100), celsiusToFahrenheit(0) 
}

func main() {
	fmt.Printf("100°C = %.2f \n", celsiusToFahrenheit(100))
	fmt.Printf("72°F = %.2f \n", fahrenheitToCelsius(72))
	b, f := boilingAndFreezing()
	fmt.Printf("Boiling: %v°F, Freezing: %v°F \n", b, f)
	fmt.Printf("72°F as whole number Celsius: %.0f \n", fahrenheitToCelsius(72))

	var f64 float64
	fmt.Printf("Zero value of float64: %f \n", f64)
}