package main

import "fmt"

func main() {
	type line struct {
		x1 int
		y1 int
		x2 int
		y2 int
	}
	// l1 := line{2, 3, 4, 6}
	// l2 := line{3, 3, 8, 8}
	l1 := line{
		x1: 2,
		y1: 3,
	}
	fmt.Println(l1.x1)
	l1.x1 = 10

	// var numberA int = 4
	// var numberB *int = &numberA

	// fmt.Println("numberA (value)   :", numberA)  // 4
	// fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	// fmt.Println("numberB (value)   :", *numberB) // 4
	// fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)
}