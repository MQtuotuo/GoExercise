package main

import "fmt"

//var language string

type Framework string
type Language string
type Minute int
type Hour int

func main() {
	language := Language("Java")
	framework := Framework("Spring 5")

	fmt.Println("Language to learn:", language)
	fmt.Println("Framework to learn:", framework)

	language = language + " - Ruby"

	fmt.Println(language)
	print(string(language))

	minutes := Minute(70)
	hour := Hour(10)

	if minutes > 60 {
		fmt.Println("Minutes is greater than 60")
	}

	if hour < 15 {
		fmt.Println("Hhours is greater than 15")
	}
}

func print(value string) {
	fmt.Println("Value:", value)
}
