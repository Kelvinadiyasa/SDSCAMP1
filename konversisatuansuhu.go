package main

import (
	"fmt"
)

// Fungsi untuk mengkonversi suhu dari Celsius ke Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9/5) + 32
}

// Fungsi untuk mengkonversi suhu dari Celsius ke Kelvin
func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

// Fungsi untuk mengkonversi suhu dari Fahrenheit ke Celsius
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5/9
}

// Fungsi untuk mengkonversi suhu dari Fahrenheit ke Kelvin
func fahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32)*5/9 + 273.15
}

// Fungsi untuk mengkonversi suhu dari Kelvin ke Celsius
func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

// Fungsi untuk mengkonversi suhu dari Kelvin ke Fahrenheit
func kelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273.15)*9/5 + 32
}

func main() {
	// Contoh penggunaan fungsi-fungsi konversi suhu
	celsius := 25.0
	fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", celsius, celsiusToFahrenheit(celsius))
	fmt.Printf("%.2f Celsius = %.2f Kelvin\n", celsius, celsiusToKelvin(celsius))

	fahrenheit := 77.0
	fmt.Printf("%.2f Fahrenheit = %.2f Celsius\n", fahrenheit, fahrenheitToCelsius(fahrenheit))
	fmt.Printf("%.2f Fahrenheit = %.2f Kelvin\n", fahrenheit, fahrenheitToKelvin(fahrenheit))

	kelvin := 298.15
	fmt.Printf("%.2f Kelvin = %.2f Celsius\n", kelvin, kelvinToCelsius(kelvin))
	fmt.Printf("%.2f Kelvin = %.2f Fahrenheit\n", kelvin, kelvinToFahrenheit(kelvin))
}