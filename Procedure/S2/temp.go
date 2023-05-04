package main

import "fmt"

func cCtoRFK(cel float64) {
	var R, F float64
	var K float64

	R = (4 * cel) / 5
	F = ((9 * cel) / 5) + 32
	K = cel + 273.15
	fmt.Printf("%vR %vF %.2fK\n", R, F, K)
}

func cCtoRFK1(cel float64, R, F, K *float64) {
	*R = (4 * cel) / 5
	*F = ((9 * cel) / 5) + 32
	*K = cel + 273.15
}

func main() {
	var c, Re, Fa, Ke float64
	fmt.Scan(&c)

	// procedure with input only
	fmt.Println("output 1")
	cCtoRFK(c)

	// procedure with input/output
	fmt.Println("output 2")
	cCtoRFK1(c, &Re, &Fa, &Ke)
	fmt.Printf("%vR %vF %.2fK\n", Re, Fa, Ke)
}
