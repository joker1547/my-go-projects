package main

import "fmt"

func main() {
	foo()

	n, err := fmt.Println("Hello Santhose", 42, true)
	fmt.Println(n, err)

	m, _ := fmt.Println("No error capture")

	fmt.Println(m)

	n, _ = fmt.Println("Notice ':' is used when declaring a new variable only")

}

func foo() {
	fmt.Println("Hello from func FOO")
}
