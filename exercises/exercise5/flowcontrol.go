package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Value:", i)
	}

	i := 0
	for i < 5 {
		fmt.Println("Value:", i)
		i++
	}

	for {
		importantValue := importantFunction()
		if importantValue == 10 {
			fmt.Println(importantValue)
			break
		}
	}

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	number := 10
	amount := 15

	if number >= 5 && amount <= 20 {
		fmt.Println("This will be printed!")
	}

	if number >= 15 || amount <= 20 {
		fmt.Println("This will be printed!")
	}

	if number > 5 {
		fmt.Println("Will not be printed!")
	} else if number < 7 {
		fmt.Println("Will not be printed!")
	} else if number < 9 {
		fmt.Println("Will not be printed!")
	} else {
		fmt.Println("Message Default!")
	}

	switch number {
	case 5:
		fmt.Println("The number is 5")
	case 10:
		fmt.Println("The number is 10")
	case 15:
		fmt.Println("This will not be checked!")
	}

	switch number {
	case 3, 8, 15:
		fmt.Println("The number is in the first range!")
	case 9, 10, 20:
		fmt.Println("The number is in the second range!")
	case 16, 21, 25:
		fmt.Println("The number is in the third range!")
	}

	switch {
	case number > 2 && number <= 8:
		fmt.Println("The number belongs to the first conditional")
	case number > 5 && number <= 10:
		fmt.Println("The number belongs to the second conditional")
	case number > 15 && number <= 20:
		fmt.Println("The number belongs to the third conditional")
	}

}

func importantFunction() int {
	return 10
}
