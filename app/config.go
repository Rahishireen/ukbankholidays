package app

import "fmt"

var (
	// Default port for the server.
	// This variable holds the default port number used when starting the server.
	port string = "8080"

	// The full address to bind the server to, including the port.
	addr string = fmt.Sprintf(":%v", port)
)
