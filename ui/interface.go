package ui

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func SubMenu(title string, isDoubleInput bool) (float64, float64, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%v\n\n", title)
	var firstNum, secondNum float64
	num1, err := getInput("First number: ", reader)
	if err != nil {
		return 0, 0, err
	}

	firstNum, err = strconv.ParseFloat(num1, 64)
	if err != nil {
		err = errors.New("input must be a number")
		return 0, 0, err
	}

	if isDoubleInput {
		num2, err := getInput("Second number [exponent if Power]: ", reader)
		if err != nil {
			return 0, 0, err
		}

		secondNum, err = strconv.ParseFloat(num2, 64)
		if err != nil {
			err = errors.New("input must be number")
			return 0, 0, err
		}
	}

	return firstNum, secondNum, nil
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		err = errors.New("reading input interrupted")
		return "", err
	}

	return strings.TrimSpace(input), err
}
