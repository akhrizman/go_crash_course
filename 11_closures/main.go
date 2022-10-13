package main

import "fmt"

func adder() func(int) int {
	sum := 0
	fmt.Println("anonymous function sum: %v", sum)
	return func(x int) int {
		sum += x
		fmt.Println("anonymous function number to add: %v", x)
		return sum
	}
}

func main() {
	mySum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(mySum(i))
	}
}
