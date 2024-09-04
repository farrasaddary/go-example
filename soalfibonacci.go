package main

import "fmt"

func main() {
	var in, f0, f1, next, i int
	f0 = 0
	f1 = 1
	next = 0
	fmt.Print("Masukan : ")
	fmt.Scan(&in)
	fmt.Print("Deret fibonacci : ")
	i = 0
	if i == 0 {
		fmt.Print(" ", f0, " ")
		i++
	}
	if i == 1 && i <= in {
		fmt.Print(f1, " ")
		i++
	}
	for i <= in && i < 2 {
		next = f0 + f1
		f0 = f1
		f1 = next
		fmt.Print(next, " ")
		i++
	}
	fmt.Print("\nDeret Fibonacci Selesai...")
}
