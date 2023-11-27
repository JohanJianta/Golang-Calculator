package ui

import (
	"Golang-Calculator/calculator"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Program kalkulator di terminal
func CalcProgram() {
	for {
		operation, err := menu()
		// Jika terjadi error, lanjut ke loop selanjutnya
		if checkError(err) {
			continue
		}

		var result float64

		switch operation {
		case "1":
			num1, num2, err := subMenu("Addition", true)
			if checkError(err) {
				continue
			}

			result = calculator.Addition(num1, num2)
		case "2":
			num1, num2, err := subMenu("Subtraction", true)
			if checkError(err) {
				continue
			}

			result = calculator.Subtraction(num1, num2)
		case "3":
			num1, num2, err := subMenu("Multiplication", true)
			if checkError(err) {
				continue
			}

			result = calculator.Multiplication(num1, num2)
		case "4":
			num1, num2, err := subMenu("Division", true)
			if checkError(err) {
				continue
			}

			result, err = calculator.Division(num1, num2)
			if checkError(err) {
				continue
			}
		case "5":
			num1, num2, err := subMenu("Power", true)
			if checkError(err) {
				continue
			}

			result = calculator.Power(num1, num2)
		case "6":
			// _ berarti variabel tidak akan digunakan
			num1, _, err := subMenu("Square Root", false)
			if checkError(err) {
				continue
			}

			result, err = calculator.SquareRoot(num1)
			if checkError(err) {
				continue
			}
		case "0":
			fmt.Println("Goodbye")
			return
		default:
			fmt.Println("Invalid operation")
			continue
		}

		fmt.Printf("\nResult: %v\n", result)
	}
}

// Menu utama
func menu() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- CALCULATOR ---")
	fmt.Println("Choose Operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Power")
	fmt.Println("6. Square Root")
	fmt.Println("0. Exit")
	operation, err := getInput("\n-> ", reader)
	if err != nil {
		return "", err
	}

	fmt.Println("\n---")

	return operation, nil
}

// Menu operasi yang dipilih user
func subMenu(title string, isDoubleInput bool) (float64, float64, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%v\n\n", title)

	var firstNum, secondNum float64
	num1, err := getInput("First number: ", reader)
	if err != nil {
		return 0, 0, err
	}

	// Coba parsing string inputan ke float
	firstNum, err = strconv.ParseFloat(num1, 64)
	if err != nil {
		err = errors.New("input must be a number")
		return 0, 0, err
	}

	// Tampilkan inputan kedua apabila bool true (false hanya untuk akar kuadrat)
	if isDoubleInput {
		num2, err := getInput("Second number [exponent if Power]: ", reader)
		if err != nil {
			return 0, 0, err
		}

		// Coba parsing string inputan ke float
		secondNum, err = strconv.ParseFloat(num2, 64)
		if err != nil {
			err = errors.New("input must be number")
			return 0, 0, err
		}
	}

	// Kembalikan angka pertama dan kedua apabila berhasil parsing
	return firstNum, secondNum, nil
}

// Tampilkan pesan dan baca inputan user
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	// Apabila terjadi error ketika menginput (seperti ^C)
	if err != nil {
		err = errors.New("reading input interrupted")
		return "", err
	}

	// Hapus spasi di depan dan belakang inputan
	return strings.TrimSpace(input), err
}

// Print pesan error apabila terjadi kesalahan
func checkError(err error) bool {
	if err != nil {
		fmt.Printf("\n[ERROR] %v\n", err)
		return true
	}

	return false
}
