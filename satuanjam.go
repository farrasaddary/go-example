package main

import "fmt"

func main() {
	var awal, detik, jam, menit int
	fmt.Print("Masukan : ")
	fmt.Scan(&awal)

	if awal/3600 != 0 {
		jam = awal / 3600
		awal = awal - (jam * 3600)
	}
	if awal/60 != 0 {
		menit = awal / 60
		awal = awal - (menit * 60)
	}
	detik = awal

	fmt.Print(jam, " Jam ", menit, " Menit ", detik, " Detik")
}
