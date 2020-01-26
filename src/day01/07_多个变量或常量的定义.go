package main

import "fmt"

func main() {
	var (
		a = 1
		b = 2.0
	)
	a, b = 10, 3.14
	fmt.Printf("a: %d, b: %f\n", a, b)

	const (
		i = 10
		j = 3.14
	)
	fmt.Printf("i: %d, j: %f\n", i, j)
}
