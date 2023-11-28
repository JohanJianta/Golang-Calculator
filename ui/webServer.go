// File webServer.go berisikan server untuk menjalankan program kalkulator melalui peramban web
// [Note] ini hanya testing semata :)

package ui

import (
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
			result, _ = calcFunc[0].Function(num1, num2)
		case "subtraction":
			result, _ = calcFunc[1].Function(num1, num2)
		case "multiplication":
			result, _ = calcFunc[2].Function(num1, num2)
		case "division":
			result, err = calcFunc[3].Function(num1, num2)
		case "power":
			result, _ = calcFunc[4].Function(num1, num2)
		case "squareRoot":
			result, err = calcFunc[5].Function(num1, num2)
		case "sine":
			result, err = calcFunc[6].Function(num1, num2)
		case "cosine":
			result, err = calcFunc[7].Function(num1, num2)
		case "tangent":
			result, err = calcFunc[8].Function(num1, num2)
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

	// Jalankan server di localhost:8080
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
