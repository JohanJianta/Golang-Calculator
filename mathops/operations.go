// File operations.go berisikan 4 function operasi matematika dasar
// Pertambahan, pengurangan, perkalian, dan pembagian

package mathops

import "errors"

// Penjumlahan
func Add(a, b float64) float64 {
	return a + b
}

// Pengurangan
func Subtract(a, b float64) float64 {
	return a - b
}

// Perkalian
func Multiply(a, b float64) float64 {
	return a * b
}

// Pembagian
func Divide(a, b float64) (float64, error) {
	// apabila mencoba membagi dengan 0
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
