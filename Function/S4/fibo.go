package main

import "fmt"

func fib(angka int) int {
	var i, then, now, tot, num int
	then, tot, num, now = 0, 0, 0, 1
	for i <= angka {
		if i == 1 || i == 0 {
			tot += then
		} else if i == 2 {
			tot += now
		} else {
			num = then + now
			then = now
			now = num
			tot += num
		}
		i++
	}
	return tot
}

func main() {
	var in int

	fmt.Scan(&in)
	fmt.Println(fib(in))
}
