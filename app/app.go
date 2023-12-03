package app

import (
	"log"
	"net/http"
	"ukbankholidays/utils"
)

// StartServer initializes and starts the application's HTTP server
func StartServer() {
	//sets up the necessary routes
	router := setupRoutes()
	utils.Logger.Debug("Routes are set up successfully.")

	//launches the server to handle incoming HTTP requests on port 8080
	// Note: The function exits with a fatal log message if there's an issue
	// starting the server.
	utils.Logger.Infof("Server started on :%v", port)
	log.Fatal(http.ListenAndServe(addr, router))
}
