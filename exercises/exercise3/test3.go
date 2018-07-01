package main

import "fmt"

func main() {
	sumAndPrint(3, 6)
	value := sum(9, 7)
	fmt.Println("Total value:", value)

	total := sum2(8, 12)

	fmt.Println("Total value:", total)
}

func sumAndPrint(first int, second int) {
	sum := first + second

	fmt.Println(sum)

}

func sum(first, second int) int {
	totalValue := first + second

	return totalValue
}

func sum2(first, second int) (amount int) {
	amount = first + second

	return
}
