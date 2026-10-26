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
}