package main

import (
	"fmt"
)

// Fungsi untuk menghitung frekuensi dari periode
func hitungFrekuensi(periode float64) float64 {
	return 1 / periode
}

// Fungsi untuk menghitung periode dari frekuensi
func hitungPeriode(frekuensi float64) float64 {
	return 1 / frekuensi
}

func main() {
	// Memasukkan data
	var periode float64
	var frekuensi float64

	fmt.Print("Masukkan periode (s): ")
	fmt.Scan(&periode)

	fmt.Print("Masukkan frekuensi (Hz): ")
	fmt.Scan(&frekuensi)

	// Menghitung frekuensi dan periode
	frekuensiDariPeriode := hitungFrekuensi(periode)
	periodeDariFrekuensi := hitungPeriode(frekuensi)

	// Menampilkan hasil
	fmt.Printf("Frekuensi dari periode %.2f detik adalah %.2f Hz\n", periode, frekuensiDariPeriode)
	fmt.Printf("Periode dari frekuensi %.2f Hz adalah %.2f detik\n", frekuensi, periodeDariFrekuensi)
}