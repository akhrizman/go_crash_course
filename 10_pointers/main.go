package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Printf("value: %v, type: %T\n", a, a)
	fmt.Printf("value: %v, type: %T\n", b, b)

	// Use * to read val from address
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)

	// Everything in Go is Pass-By-Value

}
