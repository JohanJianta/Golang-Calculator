package calculator

import (
	"Golang-Calculator/mathops"
	"errors"
	"math"
)

// Pertambahan
func Addition(a, b float64) (float64, error) {
	return roundIfDecimalIsZero(mathops.Add(a, b)), nil
}

// Pengurangan
func Subtraction(a, b float64) (float64, error) {
	return roundIfDecimalIsZero(mathops.Subtract(a, b)), nil
}

// Perkalian
func Multiplication(a, b float64) (float64, error) {
	return roundIfDecimalIsZero(mathops.Multiply(a, b)), nil
}

// Pembagian
func Division(a, b float64) (float64, error) {
	result, err := mathops.Divide(a, b)
	if err != nil {
		return 0, err
	}
	return roundIfDecimalIsZero(result), nil
}

// Perpangkatan
func Power(base, exponent float64) (float64, error) {
	return roundIfDecimalIsZero(math.Pow(base, exponent)), nil
}

// Akar kuadrat
func SquareRoot(number float64) (float64, error) {
	// Pastikan angka tidak negatif
	if number < 0 {
		return 0, errors.New("cannot calculate square root of a negative number")
	}

	// Gunakan metode Newton untuk memperkirakan hasil akar kuadrat [yah begitulah]
	guess := number / 2.0
	epsilon := 1e-15 // tingkat keakuratan perkiraan

	for math.Abs(guess*guess-number) > epsilon {
		guess = 0.5 * (guess + number/guess)
	}

	return roundIfDecimalIsZero(guess), nil
}

func roundIfDecimalIsZero(f float64) float64 {
	// Bulatkan angka apabila decimal bernilai 0
	if math.Mod(f, 1) == 0 {
		return math.Round(f)
	}

	return f
}
