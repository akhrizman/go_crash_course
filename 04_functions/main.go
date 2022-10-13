package main

import (
	"fmt"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func getSumImplicit(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Alex"))
	fmt.Println(getSum(3, 4))
	fmt.Println(getSumImplicit(3, 4))
}
