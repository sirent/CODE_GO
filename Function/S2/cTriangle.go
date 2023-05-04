package main

import "fmt"

func checkTriangle(s1, s2, s3 int) string {
	if s1+s2 > s3 && s1+s3 > s2 && s2+s3 > s1 {
		return "segitiga"
	} else {
		return "bukan segitiga"
	}
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)
	fmt.Println(checkTriangle(a, b, c))
}
