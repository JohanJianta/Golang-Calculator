package main

import (
	"Golang-Calculator/calculator"
	"Golang-Calculator/ui"
	"fmt"
)

func main() {
	calcProgram()
}

func calcProgram() {
	for {
		operation, err := ui.Menu()
		if checkError(err) {
			continue
		}

		var result float64
		switch operation {
		case "1":
			num1, num2, err := ui.SubMenu("Addition", true)
			if checkError(err) {
				continue
			}

			result, err = calculator.Addition(num1, num2)
			if checkError(err) {
				continue
			}
		case "2":
			num1, num2, err := ui.SubMenu("Subtraction", true)
			if checkError(err) {
				continue
			}

			result, err = calculator.Subtraction(num1, num2)
			if checkError(err) {
				continue
			}
		case "3":
			num1, num2, err := ui.SubMenu("Multiplication", true)
			if checkError(err) {
				continue
			}

			result, err = calculator.Multiplication(num1, num2)
			if checkError(err) {
				continue
			}
		case "4":
			num1, num2, err := ui.SubMenu("Division", true)
			if checkError(err) {
				continue
			}

			result, err = calculator.Division(num1, num2)
			if checkError(err) {
				continue
			}
		case "5":
			num1, num2, err := ui.SubMenu("Power", true)
			if checkError(err) {
				continue
			}

			result, err = calculator.Power(num1, num2)
			if checkError(err) {
				continue
			}
		case "6":
			num1, _, err := ui.SubMenu("Square Root", false)
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

func checkError(err error) bool {
	if err != nil {
		fmt.Printf("\n[ERROR] %v\n", err)
		return true
	}

	return false
}
