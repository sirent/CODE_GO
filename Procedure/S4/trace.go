package main

import "fmt"

var a int

func procBtA(b, c int) {
	b = b + c
	c = a + b + c
}

func procBtB(b int, c *int) {
	b = b + *c
	*c = a + b + *c
}

func procBtC(b *int, c int) {
	*b = *b + c
	c = a + *b + c
}

func procBtD(b, c *int) {
	*b = *b + *c
	*c = a + *b + *c
}

func main() {
	a = 10

	// a = 10
	fmt.Println("output a")
	procBtA(a, a)
	fmt.Println(a)
	fmt.Println("")

	// b = 40
	fmt.Println("output b")
	procBtB(a, &a)
	fmt.Println(a)
	fmt.Println("")

	// c = 20
	fmt.Println("output c")
	procBtC(&a, a)
	fmt.Println(a)
	fmt.Println("")

	// d = 60
	fmt.Println("output d")
	procBtD(&a, &a)
	fmt.Println(a)
}
