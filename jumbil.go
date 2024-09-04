package main

import "fmt"

func main() {
	var a, akhir int
	fmt.Print("Masukkan")
	fmt.Scan(&a)
	for a/10 != 0 || a > 0 {
		akhir = akhir + a%10
		a = a / 10
	}
	fmt.Print("Hasil : ", akhir)
}
