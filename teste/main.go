package main

import "strings"

func main() {

	result := "A"

	result += strings.Repeat(" ", 3)

	println(result)

}
