package mathops

import (
	"fmt"
	"testing"
)

// Tes penjumlahan
func TestAddition(t *testing.T) {
	result := Add(3, 4)
	expected := 7.0
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	} else {
		fmt.Println("[CORRECT] Addition")
	}
}

// Tes pengurangan
func TestSubtraction(t *testing.T) {
	result := Subtract(8, 4)
	expected := 4.0
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	} else {
		fmt.Println("[CORRECT] Subtraction")
	}
}

// Tes perkalian
func TestMultiplication(t *testing.T) {
	result := Multiply(5, 3)
	expected := 15.0
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	} else {
		fmt.Println("[CORRECT] Multiplication")
	}
}

// Tes pembagian
func TestDivision(t *testing.T) {
	result, err := Divide(10, 2)
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
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Expected error for division by zero, but got none")
	} else {
		fmt.Println("[CORRECT] Division Error Handling")
	}
}
