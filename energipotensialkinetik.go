package main

import "fmt"


func hitungEnergiPotensial(massa float64, ketinggian float64) float64 {
	return massa * 9.8 * ketinggian
}

func hitungEnergiKinetik(massa float64, kecepatan float64) float64 {
	return 0.5 * massa * kecepatan * kecepatan
}

func main() {
	var massa float64
	var ketinggian float64
	var kecepatan float64

	fmt.Print("Masukkan massa (kg): ")
	fmt.Scan(&massa)

	fmt.Print("Masukkan ketinggian (m): ")
	fmt.Scan(&ketinggian)

	fmt.Print("Masukkan kecepatan (m/s): ")
	fmt.Scan(&kecepatan)

	energiPotensial := hitungEnergiPotensial(massa, ketinggian)
	energiKinetik := hitungEnergiKinetik(massa, kecepatan)
	fmt.Printf("Energi Potensial: %.2f J\n", energiPotensial)
	fmt.Printf("Energi Kinetik: %.2f J\n", energiKinetik)
}