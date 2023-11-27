package ui

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Menu utama
func Menu() (string, error) {
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
func SubMenu(title string, isDoubleInput bool) (float64, float64, error) {
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
