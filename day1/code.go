package main

import (
	"errors"
	"fmt"
)

//sentinel error
var (
ErrDivideByZero = errors.New("cannot divide by zero")
ErrUnsupportedOp = errors.New("unsupported operation")
)

func add(a int, b int) int {
	return a + b
	}
	func substract(a int, b int) int {
	return a - b
	}

// universal var
const PI = 3.14

//iota
const (
Summer = iota // 0
Autumn        // 1
Winter        // 2
Spring        // 3
)


	//func
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("can't divide by 0")
	}
	return a / b, nil
}

type honorificDecorator func(string) string
func createHonorificDecorator(honorific string) honorificDecorator {
	return func(name string) string {
		return honorific + ". " + name
	}
}

func main(){

	// loop, labeling, if
	// n := 5
	// here: for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if i < j {
	// 			continue here
	// 		}
	// 		fmt.Println(i, j)
	// 	}
	// }

	//const 
	const WORLD = "世界"
	fmt.Println("Hello", WORLD)
	fmt.Println("Happy", PI, "Day")

	const TRUTH = true
	fmt.Println("Go rules?", TRUTH)

	
    fmt.Println(Winter) // Outputs: 2

	//error
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	//slice
	// x := []int{1, 2, 3, 4, 5}
	// x[0] = 10
	// fmt.Println(x[1])
	// x = append(x, 3)
	// x = append(x, 4, 5)
	// fmt.Println(x)
	// fmt.Println(len(x), cap(x))

	// x := []int{1, 2, 3}
	// y := x
	// y[0] = 100
	// fmt.Println(x[0])

	// 	x := []int{1, 2, 3}
	// 	// y := x[1:3] //index 1-3 exclusive
	// 	y := x[1:] //until end
	// 	z := x[:2] //until beginning
	// 	fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(y[0])


	// 	greeting := "Hello!"
	// exclaimed := "何?!"
	greeting := "Hello!"
	gb := []byte(greeting)
	gr := []rune(greeting)
	fmt.Println(len(gb))
	fmt.Println(len(gr))
	exclaimed := "何?!"
	eb := []byte(exclaimed)
	er := []rune(exclaimed)
	fmt.Println(len(eb))
	fmt.Println(len(er))

	myMap := map[string]int{}
	myMap["one"] = 1
	fmt.Println(myMap["one"])

	//for range
	//iterate map
	greetings := map[string]string{
		"English": "Hello",
		"French": "Bonjour",
		"Japanese": "こんにちは",
	}
	for key, value := range greetings {
		fmt.Println(key, value)
	}
	//iterate string
	greeting2 := "こんにちは"
	for _, ch := range greeting2 {
	fmt.Printf("%c\n", ch)
	}

	//function as values
	
	ops := map[string] func(int, int) int {
	"+": add,
	"-": substract,
	}
	var opr1, opr2 int
	var op string
	fmt.Scan(&opr1, &op, &opr2)
	fmt.Println(ops[op](opr1, opr2))
	addMr := createHonorificDecorator("Mr")
	addMs := createHonorificDecorator("Ms")
	fmt.Println(addMr("Doe"))
	fmt.Println(addMs("Doe"))
	fmt.Println(addMr("Holmes"))
	fmt.Println(addMs("Adler"))
}

