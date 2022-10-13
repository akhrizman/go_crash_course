package main

import (
	"fmt"
)

func main() {
	//Using var
	var name string = "Brad"
	var name2 = "Brad"
	var age int = 37
	var age2 = 37
	const isCool = true
	var size = 1.7
	title := "yay"

	var one, two = "one", "two"
	three, four, five := "three", "four", 5

	fmt.Println(name, name2, age, age2, isCool, title, size)

	fmt.Printf("%T\n", age)

	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)

	fmt.Println(one, two, three, four, five)

}
