package main

import "fmt"

func pecahUang(uang int, k10, k2, k1 *int) {
	*k10 = uang / 10000
	uang = uang % 10000
	*k2 = uang / 2000
	uang = uang % 2000
	*k1 = uang / 1000
}

func main() {
	var money, tth, twth, oth int

	fmt.Scan(&money)

	pecahUang(money, &tth, &twth, &oth)
	fmt.Println(tth, "lembar 10000")
	fmt.Println(twth, "lembar 2000")
	fmt.Println(oth, "lembar 1000")
}
