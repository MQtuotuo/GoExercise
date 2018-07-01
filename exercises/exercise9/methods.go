package main

import (
	"fmt"
	"strings"
)

type Description string

func (d Description) Upper() string {
	return strings.ToUpper(string(d))
}

type VideoCourse struct {
	Minute   int
	Language string
}

func (v VideoCourse) Description() string {
	return fmt.Sprintf("Video Course for %s with a duration of %d minutes", v.Language, v.Minute)
}

func main() {
	description := Description("My Go special description")
	upper := description.Upper()

	fmt.Println(upper)

	course := VideoCourse{Language: "Go", Minute: 60}
	fmt.Println(course)

	descript := course.Description()
	fmt.Println(descript)
}
