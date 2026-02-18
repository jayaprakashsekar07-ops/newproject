package main

import "fmt"

func main() {

	// Basic types
	var a int
	var b float64
	var c string
	var d bool

	// Pointer
	var p *int

	// Slice
	var s []int

	// Map
	var m map[string]int

	m["a"] = 10

	fmt.Println("int:", a)
	fmt.Println("float:", b)
	fmt.Println("string:", c)
	fmt.Println("bool:", d)
	fmt.Println("pointer:", p)
	fmt.Println("slice:", s)
	fmt.Println("map:", m)
}
