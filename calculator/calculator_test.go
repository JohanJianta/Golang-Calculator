// File ini untuk tes function operasi kalkulator dari calculator.go
// Untuk memulai tes, buka terminal dan pindah ke lokasi file ini
// Kemudian jalankan prompt "go test"

package calculator

import (
	"fmt"
	"testing"
)

// Tes penjumlahan
func TestAddition(t *testing.T) {
	result, _ := Addition(3, 4)
	expected := 7.0
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	} else {
		fmt.Println("[CORRECT] Addition")
	}
}

// Tes pengurangan
func TestSubtraction(t *testing.T) {
	result, _ := Subtraction(8, 4)
	expected := 4.0
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	} else {
		fmt.Println("[CORRECT] Subtraction")
	}
}

// Tes perkalian
func TestMultiplication(t *testing.T) {
	result, _ := Multiplication(5, 3)
	expected := 15.0
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	} else {
		fmt.Println("[CORRECT] Multiplication")
	}
}

// Tes pembagian
func TestDivision(t *testing.T) {
	result, err := Division(10, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := 5.0
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	} else {
		fmt.Println("[CORRECT] Division")
	}

	// Tes pembagian dengan 0
	_, err = Division(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero, but got none")
	} else {
		fmt.Println("[CORRECT] Division Error Handling")
	}
}

// Tes perpangkatan
func TestPower(t *testing.T) {
	result, _ := Power(2, 3)
	expected := 8.0
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	} else {
		fmt.Println("[CORRECT] Power")
	}
}

// Tes akar kuadrat
func TestSquareRoot(t *testing.T) {
	result, _ := SquareRoot(9, 0)
	expected := 3.0
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	} else {
		fmt.Println("[CORRECT] Square Root")
	}
}
