package main

import (
	"fmt"
	"math"
)
func main() {
	const inflationRate =6.5
	var investmentAmount, years, expectedReturnRate float64
	
	fmt.Print("Please, enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Please, enter amount of years: ")
	fmt.Scan(&years)
	fmt.Print("Please, enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)



	futureValue := investmentAmount*math.Pow(1+expectedReturnRate/100,years)
	
	futureRealValue:=futureValue/ math.Pow(1+inflationRate/100,years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
