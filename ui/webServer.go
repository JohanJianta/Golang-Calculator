package ui

import (
	"Golang-Calculator/calculator"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WebServer() {
	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Serve static files from the "static" directory
	r.Static("/static", "./static/")

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/calculate", func(c *gin.Context) {
		operation := c.PostForm("operation")
		num1, _ := strconv.ParseFloat(c.PostForm("num1"), 64)
		num2, _ := strconv.ParseFloat(c.PostForm("num2"), 64)

		var result float64
		var err error

		switch operation {
		case "addition":
			result = calculator.Addition(num1, num2)
		case "subtraction":
			result = calculator.Subtraction(num1, num2)
		case "multiplication":
			result = calculator.Multiplication(num1, num2)
		case "division":
			result, err = calculator.Division(num1, num2)
		case "power":
			result = calculator.Power(num1, num2)
		case "squareRoot":
			result, err = calculator.SquareRoot(num1)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid operation"})
			return
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "result.html", gin.H{
			"operation": operation,
			"num1":      num1,
			"num2":      num2,
			"result":    result,
		})
	})

	// Run the server
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
