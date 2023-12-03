package main

import (
	"ukbankholidays/app"
	"ukbankholidays/utils"
)

// main launches the application server
func main() {
	utils.InitLog()

	utils.Logger.Debug("Starting the application...")

	app.StartServer()
}
