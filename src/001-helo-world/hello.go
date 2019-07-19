package main

import "fmt"

var z = 43 // use var to declare outside of function body
var k int  // declare there is a variable with identifier k of type int
// nil for pointers, false for booleans, 0.0 for floats and 0 for ints., "" for strings.
var variable = `this is a "test to include double quotes"` // back quotes not single quote

func main() {
	foo()

	n, err := fmt.Println("Hello Santhose", 42, true)
	fmt.Println(n, err)

	m, _ := fmt.Println("No error capture")

	fmt.Println(m)

	n, _ = fmt.Println("Notice ':' is used when declaring a new variable only")

	x := 42
	fmt.Println(x)

	fmt.Println(variable)

	fmt.Printf("variable is of type: %T\n", variable) //%T prints type of the variable

	x = 43

	y := x * x

	fmt.Println("x= ", x, "y = ", y)

	fmt.Println(z)
}

func foo() {
	fmt.Println("Hello from func FOO")
	fmt.Println("In foo z is equal to", z, "value of K = ", k)
}
