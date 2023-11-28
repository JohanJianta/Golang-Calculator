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

// Array operasi kalkulator beserta namanya
// Di buat demikin agar switch case bisa dipersingkat
var calcFunc = []calculator.CalculatorFunction{
	{Title: "Addition", Function: calculator.Addition},
	{Title: "Subtraction", Function: calculator.Subtraction},
	{Title: "Multiplication", Function: calculator.Multiplication},
	{Title: "Division", Function: calculator.Division},
	{Title: "Power", Function: calculator.Power},
	{Title: "Square Root", Function: calculator.SquareRoot},
	{Title: "Sine", Function: calculator.Sine},
	{Title: "Cosine", Function: calculator.Cosine},
	{Title: "Tangent", Function: calculator.Tangent},
}

// Program kalkulator di terminal
func CalcProgram() {
	for {
		// Tampilkan menu pilihan
		operation, err := menu()
		// Jika terjadi error, skip switch dan lanjut ke loop selanjutnya
		if printError(err) {
			continue
		}

		fmt.Println("\n---")

		switch operation {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			// Parse string ke integer
			index, _ := strconv.ParseInt(operation, 10, 64)

			// SubMenu untuk ambil inputan pertama dan kedua
			num1, num2, err := subMenu(calcFunc[index-1].Title)
			if printError(err) {
				break
			}

			// Jalankan operasi kalkulator sesuai dengan operasi yang dipilih
			result, err := calcFunc[index-1].Function(num1, num2)
			if printError(err) {
				break
			}

			// %.10f berarti tampilkan float sampai 10 desimal
			fmt.Printf("\nResult: %.10f\n", result)
		case "0":
			// Hentikan program
			fmt.Println("Goodbye")
			return
		default:
			// Infokan apabila pilihan tidak tersedia
			fmt.Println("Invalid Operation")
		}

		fmt.Println("---")
	}
}

// Menu utama
func menu() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n--- CALCULATOR ---")
	fmt.Println("Choose Operation:")
	// Loop print seluruh operasi kalkulator yang tersedia
	for index, value := range calcFunc {
		// Contoh format "1. Addition"
		fmt.Printf("%v. %v\n", (index + 1), value.Title)
	}
	fmt.Println("0. Exit")
	operation, err := getInput("\n-> ", reader)
	if err != nil {
		return "", err
	}

	return operation, nil
}

// SubMenu dari operasi yang dipilih user
func subMenu(title string) (float64, float64, error) {
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

	// Hanya tampilkan inputan kedua apabila bukan akar kuadrat
	if title != "Square Root" && title != "Sine" && title != "Cosine" && title != "Tangent" {
		num2, err := getInput("Second number [exponent if Power]: ", reader)
		if err != nil {
			return 0, 0, err
		}

		// Coba parsing string inputan ke float
		secondNum, err = strconv.ParseFloat(num2, 64)
		if err != nil {
			err = errors.New("input must be a number")
			return 0, 0, err
		}
	}

	// Kembalikan angka pertama dan kedua apabila berhasil parsing
	return firstNum, secondNum, nil
}

// Tampilkan pesan dan baca inputan user
// Parameter reader pakai pointer (*) agar tidak perlu dibuat ulang variabelnya
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
func printError(err error) bool {
	if err != nil {
		fmt.Printf("\n[ERROR] %v\n", err)
		return true
	}

	return false
}
