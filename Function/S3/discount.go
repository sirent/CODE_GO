package main

import "fmt"

func Disc(tbAwal int, m bool) float64 {
	if tbAwal > 100000 && m {
		return float64(tbAwal) * 0.1
	} else if tbAwal > 100000 && !m {
		return float64(tbAwal) * 0.05
	} else {
		return 0
	}
}

func main() {
	var tBelanjaAwal int
	var membership bool

	fmt.Scan(&tBelanjaAwal, &membership)
	fmt.Println(tBelanjaAwal - int(Disc(tBelanjaAwal, membership)))
}
