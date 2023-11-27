// Untuk menjalankan program kalkulator, buka proyek ini di terminal dan jalankan prompt "go run main.go"
// Uncomment ui.CalcProgram() untuk menjalankan program kalkulator di terminal [recommended]
// Uncomment ui.WebServer() untuk menjalankan program kalkulator di web localhost:8080

package main

import "Golang-Calculator/ui"

func main() {
	// Program terminal
	ui.CalcProgram()

	// Program web
	// ui.WebServer()
}
