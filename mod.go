package main

import "fmt"

func main() {
	var kode int
	fmt.Print("Masukkan : ")
	fmt.Scan(&kode)
	fmt.Print("Hasil mod : ", kode%100)
}
