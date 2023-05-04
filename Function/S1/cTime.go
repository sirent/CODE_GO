package main

import "fmt"

func convertTime(hour, minute, second int) int {
	var total int
	total = hour*3600 + minute*60 + second
	return total
}

func main() {
	var jam, menit, detik int
	fmt.Scan(&jam, &menit, &detik)
	fmt.Println("Hasil konversi =", convertTime(jam, menit, detik), "detik")
}
