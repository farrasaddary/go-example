package main

import "fmt"

func main() {
	var kode, discon, cashback, a, b int
	fmt.Print("Masukkan : ")
	fmt.Scan(&kode)
	discon = (kode / 10) % 100
	a = kode / 1000
	b = (kode / 10) % 10
	cashback = a + b
	if discon%2 == 0 {
		fmt.Println("Diskon?true")
	} else {
		fmt.Println("Diskon?false")
	}
	if kode%10 != 0 {
		if cashback%(kode%10) == 0 {
			fmt.Println("Cashback?true")
		} else {
			fmt.Println("Cashback?false")
		}
	} else {
		fmt.Println("Cashback?false")
	}

}
