package main

import "fmt"

func main() {
	foo()

	n, err := fmt.Println("Hello Santhose", 42, true)
	fmt.Println(n, err)

	m, _ := fmt.Println("No error capture")

	fmt.Println(m)

	n, _ = fmt.Println("Notice ':' is used when declaring a new variable only")

	x := 42
	fmt.Println(x)

	x = 43

	y := x * x

	fmt.Println("x= ", x , "y = ", y)

}

func foo() {
	fmt.Println("Hello from func FOO")
}
