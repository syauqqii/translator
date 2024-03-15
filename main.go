package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"translator/config"
	"translator/route"
)

func main() {
	config.LoadConfig()
	route.SetupRoutes()

	port := 6666
	fmt.Printf("Server is running at http://localhost:%d\n", port)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh

	fmt.Println("\nReceived Ctrl+C signal. Shutting down server...")
	os.Exit(0)
}
