package main

import "fmt"

const NMax int = 5

type angka struct {
	bulat int
}

type tabInt [NMax]angka

func BinBoolean(tab tabInt, n, x int) bool {
	var left, right, mid int
	left = 1
	right = n
	mid = (left + right) / 2
	for left <= right && tab[mid].bulat != x {
		if x < tab[mid].bulat {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return mid > 0 && tab[mid].bulat == x
}

func BinArray(tab tabInt, n, x int) int {
	var left, right, mid, found int
	left = 1
	right = n
	found = -1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x < tab[mid].bulat {
			right = mid - 1
		} else if x > tab[mid].bulat {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

func main() {
	var a tabInt
	var i, x, p int

	// INPUT ARRAY
	i = 0
	for i <= NMax-1 {
		fmt.Scan(&a[i].bulat)
		i += 1
	}

	// OUTPUT ARRAY SATU-SATU
	i = 0
	for i <= NMax-1 {
		fmt.Print(a[i].bulat, " ")
		i += 1
	}
	fmt.Println("")

	// PANGGIL FUNGSI BINARY SEARCH DENGAN OUTPUT BOOLEAN
	x = 1                         // NILAI YANG MAU DICARI DARI ARRAY
	fmt.Println(BinBoolean(a, NMax, x)) // OUTPUT TRUE ATAU FALSE SAJA

	// PANGGIL FUGNSI BINARY SEARCH DENGAN OUTPUT INDEX ARRAY
	fmt.Println(BinArray(a, NMax, x)) // OUTPUT INDEX ARRAY SAJA
	p = BinArray(a, NMax, x)
	fmt.Println("Hasil Binary Search nilai ditemukan di array ke-", p, "dengan nilai yaitu", a[p].bulat) // OUTPUT NILAI ARRAYNYA
}
