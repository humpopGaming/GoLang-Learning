package main

import "fmt"

const (
	USDToEur = 0.92
	USDToGbp = 0.79
	USDToJPY = 154.50
)

var bankName = "Go Exchange"

func main() {
	amountUSD := 100.0
	var euros = amountUSD * USDToEur
	var pounds float64 = amountUSD * USDToGbp
	yen := amountUSD * USDToJPY

	fmt.Println("Welcome to " + bankName)
	fmt.Printf("$100.00 = €%.2f\n", euros)
	fmt.Printf("$100.00 = £%.2f\n", pounds)
	fmt.Printf("$100.00 = ¥%.2f\n", yen)
	fmt.Printf("Total value in foreign currencies: %.2f\n", euros + pounds + yen)
}
