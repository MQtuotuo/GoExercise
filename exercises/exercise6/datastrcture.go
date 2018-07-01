package main

import "fmt"

func main() {
	var languages [5]string
	languages[0] = "Go"
	languages[1] = "Ruby"
	languages[2] = "Pony"
	languages[3] = "Erlang"
	languages[4] = "Java"

	fmt.Println(languages)
	fmt.Println(languages[3])
	fmt.Println("Array size:", len(languages))

	//Slices holds under the hood an Array to keep the items
	//Slices doesn't have a fixed size!

	languages_slice := make([]string, 3)
	languages_slice[0] = "Go"
	languages_slice[1] = "Ruby"
	languages_slice[2] = "Pony"

	fmt.Println(languages_slice)
	fmt.Println("Slice size:", len(languages_slice))
	languages_slice = append(languages_slice, "Erlang")
	languages_slice = append(languages_slice, "Java")
	fmt.Println("New slice size:", len(languages_slice))
	languages_slice = append(languages_slice, "JavaScript", "Python")
	fmt.Println("New slice size:", len(languages_slice))

	/**
	In other programming languages you have the same approach with Map, Hashes or Dictionary.
	A Map maps a set of keys to corresponding values. A few considerations about Map in Go:
	You can map any type of key in your Map
	All keys should have the same type in Go
	You can use the key to retrieve a value

	variable := map[type]type{}
	*/

	languages_map := map[string]int{}
	languages_map["java"] = 5
	languages_map["ruby"] = 4
	languages_map["go"] = 2

	fmt.Println("First value:", languages_map["java"])
	fmt.Println("Second value:", languages_map["ruby"])
	fmt.Println("Third value:", languages_map["go"])
	fmt.Println("Length of the map:", len(languages_map))

	for language, number := range languages_map {
		fmt.Println("Key:", language, "- Value:", number)
	}

	// We can use underscore to ignore the key in the loop
	for _, number := range languages_map {
		fmt.Println("Value:", number)
	}

	for key, _ := range languages {
		fmt.Println("Value:", key)
	}

	delete(languages_map, "ruby")
	delete(languages_map, "java")
	fmt.Println(languages_map)

}
