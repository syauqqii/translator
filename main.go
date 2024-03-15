package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"translator/config"
	"translator/route"
)

func main() {
	config.LoadConfig()
	route.SetupRoutes()

	fmt.Printf("\n $ INFO: Server is running at http://%s:%s\n", config.SERVER_APP, config.PORT_APP)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%s", config.PORT_APP), nil); err != nil {
			fmt.Printf(" ! ERROR: Error starting server: %s\n", err.Error())
		}
	}()

	<-sigCh

	os.Exit(0)
}
