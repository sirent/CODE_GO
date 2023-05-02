package main

import "fmt"

const nMax int = 2022

type student struct {
	name, id string
	gpa      float64
}

type tabMhs [nMax]student

func maksPos(A tabMhs, N, pass int) int {
	var idx, i int

	idx = pass - 1
	i = pass
	for i < N {
		if A[idx].gpa < A[i].gpa {
			idx = i
		}
		i++
	}
	return idx
}

func minPos(A tabMhs, N, pass int) int {
	var idx, i int

	idx = pass - 1
	i = pass
	for i < N {
		if A[idx].gpa > A[i].gpa {
			idx = i
		}
		i++
	}
	return idx
}

func tukar(A *tabMhs, i, j int) {
	var temp float64
	var temp1, temp2 string

	temp, temp1, temp2 = *&A[i].gpa, *&A[i].name, *&A[i].id
	*&A[i].gpa, *&A[i].name, *&A[i].id = *&A[j].gpa, *&A[j].name, *&A[j].id
	*&A[j].gpa, *&A[j].name, *&A[j].id = temp, temp1, temp2

}

func mengurutkan(A *tabMhs, N int) {
	var pass, idx int

	for pass = 1; pass <= N-1; pass++ {
		idx = minPos(*A, N, pass)
		tukar(A, idx, pass-1)
	}
}

func sortingFull(A *tabMhs, N int) {
	var pass, idx, i int
	var temp student

	for pass = 1; pass <= N-1; pass++ {
		idx = pass - 1
		for i = pass; i <= N-1; i++ {
			if A[idx].gpa > A[i].gpa {
				idx = i
			}
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
	}
}

func inputArr(A *tabMhs, N int) {
	var y int
	for y = 0; y <= N-1; y++ {
		fmt.Scan(&A[y].id, &A[y].name, &A[y].gpa)
	}
}

func outputArr(A *tabMhs, N int) {
	var y int
	for y = 0; y <= N-1; y++ {
		fmt.Printf("%v %s %.2f\n", A[y].id, A[y].name, A[y].gpa)
	}
}

func main() {
	var X tabMhs
	var max int

	fmt.Print("Input jumlah data:")
	fmt.Scan(&max)

	fmt.Print("Input data dalam array:")
	inputArr(&X, max)

	// fmt.Println("Pengurutan 1")
	// mengurutkan(&X, max)

	fmt.Println("Pengurutan 2")
	sortingFull(&X, max)

	fmt.Println("After sort")
	outputArr(&X, max)

}
