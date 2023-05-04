package main

import "fmt"

func mengurutkan(a, b *int) {
	var temp int = *a
	*a = *b
	*b = temp
}

func main() {
	var a1, a2 int

	fmt.Scan(&a1, &a2)
	for a1 != a2 {
		if a1 > a2 {
			mengurutkan(&a1, &a2)
			fmt.Println(a1, a2)
		} else {
			fmt.Println(a1, a2)
		}
		fmt.Scan(&a1, &a2)
	}
}
