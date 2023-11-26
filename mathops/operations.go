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
	if b == 0 {
		// apabila mencoba membagi dengan 0
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	return a / b, nil
}
