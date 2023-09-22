package main

import "fmt"

func hitungLuasPersegi(sisi float64) float64 {
	return sisi * sisi
}

func main() {
	sisi := 5.0 // Panjang sisi persegi
	luas := hitungLuasPersegi(sisi)
	fmt.Printf("Luas persegi dengan sisi %.2f adalah %.2f\n", sisi, luas)
}
