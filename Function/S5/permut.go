package main

import "fmt"

func fact(a int) int {
	var tot int
	tot = 1
	for a >= 1 {
		tot *= a
		a--
	}
	return tot
}

func perm(b, c int) int {
	var hsl int = fact(b) / fact(b-c)
	return hsl
}

func main() {
	var x, y int

	fmt.Scan(&x, &y)
	if x >= y {
		fmt.Println(fact(x), fact(y), perm(x, y))
	} else {
		fmt.Println(fact(x), fact(y), perm(y, x))
	}
}
