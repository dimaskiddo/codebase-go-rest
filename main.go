package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	svc "github.com/dimaskiddo/codebase-go-rest/service"
)

// Main Server Variable
var mainServer *svc.Server

// Init Function
func init() {
	// Initialize service
	svc.Initialize()

	// Initialize Routes
	initRoutes()

	// Initialize Server
	mainServer = svc.NewServer(svc.RouterHandler)
}

// Main Function
func main() {
	// Starting Server
	mainServer.Start()

	// Make Channel to Catch OS Signal
	osSignal := make(chan os.Signal, 1)

	// Catch OS Signal from Channel
	signal.Notify(osSignal, os.Interrupt)
	signal.Notify(osSignal, syscall.SIGTERM)

	// Return OS Signal as Exit Code
	<-osSignal

	// Termination Symbol Log Line
	fmt.Println("")

	// Stopping Server
	defer mainServer.Stop()

	// Close Any Database Connections
	if len(svc.Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(svc.Config.GetString("DB_DRIVER")) {
		case "mysql":
			defer svc.MySQL.Close()
		case "mongo":
			defer svc.MongoSession.Close()
		}
	}

	// Close Any Cache Connections
	if len(svc.Config.GetString("CACHE_DRIVER")) != 0 {
		switch strings.ToLower(svc.Config.GetString("CACHE_DRIVER")) {
		case "redis":
			defer svc.Redis.Close()
		}
	}
}
