// File calculator.go berisikan function operasi kalkulator
// Pertambahan, pengurangan, perkalian, pembagian, perpangkatan, dan akar kuadrat

package calculator

import (
	"Golang-Calculator/mathops"
	"errors"
	"math"
)

// Objek untuk representasi operasi kalkulator
type CalculatorFunction struct {
	Title    string
	Function func(float64, float64) (float64, error)
}

// Function yang errornya langsung return nil dibuat demikian agar struktur function nya
// sesuai dengan jenis yang ada di dalam struct CalculatorFunction

// Pertambahan
func Addition(a, b float64) (float64, error) {
	return mathops.Add(a, b), nil
}

// Pengurangan
func Subtraction(a, b float64) (float64, error) {
	return mathops.Subtract(a, b), nil
}

// Perkalian
func Multiplication(a, b float64) (float64, error) {
	return mathops.Multiply(a, b), nil
}

// Pembagian
func Division(a, b float64) (float64, error) {
	return mathops.Divide(a, b)
}

// Perpangkatan
func Power(base, exponent float64) (float64, error) {
	return math.Pow(base, exponent), nil
}

// Akar kuadrat
// _ sengaja ditambah agar memenuhi syarat dari function dalam CalculatorFunction
func SquareRoot(number, _ float64) (float64, error) {
	// Pastikan angka tidak negatif
	if number < 0 {
		return 0, errors.New("cannot calculate square root of a negative number")
	}

	return math.Sqrt(number), nil
}
