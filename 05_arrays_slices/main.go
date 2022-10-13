package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	// Declare and assign
	pieArr := [2]string{"Apple", "Cherry"}

	pieSlice := []string{"Apple", "Blueberry", "Strawberry, Chocolate"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	fmt.Println(pieArr)
	fmt.Println(pieArr[1])

	fmt.Println(pieSlice)
	fmt.Println(pieSlice[1])

	fmt.Println(len(pieSlice))
	fmt.Println(pieSlice[1:3])

}
