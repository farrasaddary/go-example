package main

import "fmt"

func main() {
	var a, b, c int
	var nilai bool
	nilai = true
	fmt.Print("Bilangan Bulat : ")
	fmt.Scan(&a)
	for a/10 != 0 && nilai != false {
		b = a % 10
		c = (a / 10) % 10
		if (b-c) == 1 || (b-c) == -1 {
			nilai = true
		} else {
			nilai = false
		}
		a = a / 10
		//fmt.Print("Jawabanbug ? ", b-c, " ")
	}
	fmt.Println("\nJawaban ? ", nilai)
}
