package main

import "fmt"

func main() {
	length := 1.2
	width := 2
	newWidth := float64(width)

	fmt.Println("Area is", length*newWidth)
	fmt.Println("length > width?", length > newWidth)
}
