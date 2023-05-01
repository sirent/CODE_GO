package main

import "fmt"

const NMax int = 5

type angka struct {
	bulat int
}

type tabInt [NMax]angka

func SeqBoolean(T tabInt, N int, X int) bool {
	var ketemu bool
	var k int
	ketemu = false
	k = 0
	for !ketemu && k < N {
		ketemu = T[k].bulat == X
		k += 1
	}
	return ketemu
}

func SeqIndex(T tabInt, N int, X int) int {
	var ketemu int
	var k int
	ketemu = -1
	k = 0
	for ketemu == -1 && k < N {
		if T[k].bulat == X {
			ketemu = k
		}
		k += 1
	}
	return ketemu
}

func SeqBooleanIndex(T tabInt, N int, X int) int {
	var ketemu bool
	var k int
	var z int
	ketemu = false
	k = 0
	for !ketemu && k < N {
		ketemu = T[k].bulat == X
		z = k
		k += 1
	}
	return z
}

func main() {
	var a tabInt
	var i, x, q, p int

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

	// OUTPUT ARRAY SELURUHNYA DENGAN []
	fmt.Print("\n")
	fmt.Println(a)

	// PANGGIL FUNGSI SEQUENTIAL DENGAN OUTPUT BOOLEAN
	x = 1                               // NILAI YANG MAU DICARI DARI ARRAY
	fmt.Println(SeqBoolean(a, NMax, x)) // OUTPUT HANYA TRUE ATAU FALSE

	// PANGGIL FUNGSI SEQUENTIAL DENGAN OUTPUT INDEX
	fmt.Println(SeqIndex(a, NMax, x)) // OUTPUT INDEXNYA SAJA
	q = SeqIndex(a, NMax, x)
	fmt.Println("Hasil Sequential Search nilai ditemukan di array ke-", q, "dengan nilai yaitu", a[q].bulat) // OUTPUT NILAI ARRAYNYA

	// PANGGIL FUGNSI SEQUENTIAL DENGAN OUTPUT INDEX TAPI PENCARIAN DENGAN BOOLEAN
	fmt.Println(SeqBooleanIndex(a, NMax, x)) // OUTPUT INDEXNYA SAJA
	p = SeqBooleanIndex(a, NMax, x)
	fmt.Println("Hasil Sequential Search nilai ditemukan di array ke-", p, "dengan nilai yaitu", a[p].bulat) // OUTPUT NILAI ARRAYNYA
}
