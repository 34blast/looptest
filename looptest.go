package main

import (
	"fmt"
)

func main() {
	fmt.Println("looptest.main() : starting exectuion")
	fmt.Println()

	/* create a slice */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("looptest.main() : test slice looping")
	/* print the numbers */
	for i := range numbers {
		fmt.Println("Slice item", i, "is", numbers[i])
	}
	fmt.Println()

	fmt.Println("looptest.main() : test map looping")
	/* create a map*/
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	fmt.Println()

	/* print map using key-value*/
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}

	fmt.Println()
	fmt.Println("looptest.main() : exectuion completed")

}
