package main
import "fmt"

func hitungLuasSegitiga(alas, tinggi float64) float64 {
	return 0.5 * alas * tinggi
}

func main() {
	alas := 6.0   
	tinggi := 8.0 
	luas := hitungLuasSegitiga(alas, tinggi)
	fmt.Printf("Luas segitiga dengan alas %.2f dan tinggi %.2f adalah %.2f\n", alas, tinggi, luas)
}