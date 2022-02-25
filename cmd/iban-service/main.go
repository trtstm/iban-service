/*
Entry point of our program. It is responsible for setting up all the different parts.
*/

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/trtstm/iban-service/httpserver"
	"github.com/trtstm/iban-service/iban"
)

var (
	// The address we will listen on.
	apiAddress = envString("API_ADDRESS", ":80")
)

func main() {
	// Channel that captures interupts.
	signalC := make(chan os.Signal, 1)
	// Capture interrupts.
	signal.Notify(signalC, syscall.SIGINT, syscall.SIGTERM)

	// Set up dependencies.
	logger := log.Default()
	ibanService := iban.NewIBANService()

	httpServer := httpserver.NewServer(apiAddress, ibanService, logger)
	go func() {
		logger.Printf("starting server on %s\n", apiAddress)
		httpServer.Open()
	}()

	// Wait to shutdown.
	<-signalC
	logger.Printf("shutting down")
	httpServer.Close()
}

// envString tries to get an env variable and otherwise falls back to fallback.
func envString(key string, fallback string) string {
	value, ok := syscall.Getenv(key)
	if !ok {
		return fallback
	}
	return value
}
