package calculator

import (
	"Golang-Calculator/mathops"
	"errors"
	"math"
)

// Pertambahan
func Addition(a, b float64) float64 {
	return roundIfDecimalIsZero(mathops.Add(a, b))
}

// Pengurangan
func Subtraction(a, b float64) float64 {
	return roundIfDecimalIsZero(mathops.Subtract(a, b))
}

// Perkalian
func Multiplication(a, b float64) float64 {
	return roundIfDecimalIsZero(mathops.Multiply(a, b))
}

// Pembagian
func Division(a, b float64) (float64, error) {
	return mathops.Divide(a, b)
}

// Perpangkatan
func Power(base, exponent float64) float64 {
	return roundIfDecimalIsZero(math.Pow(base, exponent))
}

// Akar kuadrat
func SquareRoot(number float64) (float64, error) {
	// Pastikan angka tidak negatif
	if number < 0 {
		return 0, errors.New("cannot calculate square root of a negative number")
	}

	// Gunakan metode Newton untuk memperkirakan hasil akar kuadrat []
	// Set nilai perkiraan awal sebagai nilai setengah dari angka
	guess := number / 2.0

	// Notasi epsilon merupakan representasi nilai kecil yang digunakan untuk membandingkan nilai float
	// Di sini epsilon merepresentasikan 1 x 10^-15 atau 0.000000000000001
	// [Note] tidak semua nilai real bisa direpresentasikan oleh komputer, makanya dibutuhkan epsilon
	// sebagai jangkauan toleransi kesalahan
	epsilon := 1e-15

	// Rumus: Iterasi Newton-Raphson
	// loop cari nilai presisi dari akar kuadrat hingga mencapai batas toleransi
	for math.Abs(guess*guess-number) > epsilon {
		guess = 0.5 * (guess + number/guess)
	}

	return roundIfDecimalIsZero(guess), nil
}

func roundIfDecimalIsZero(f float64) float64 {
	// Tetapkan jangkauan desimal yang akan dicek [1 x 10^-9]
	epsilon := 1e-9

	// Apabila sampai 9 angka desimal tidak ditemukan angka selain 0
	// maka nilai float akan langsung dibulatkan tanpa desimal
	if math.Mod(f, 1) < epsilon {
		return math.Round(f)
	}

	// Jika terdapat angka selain 0, kembalikan utuh nilai float-nya
	return f
}
