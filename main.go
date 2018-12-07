package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dimaskiddo/frame-go/utils"
)

// Main Server Variable
var mainServer *utils.Server

// Main Init Function
func init() {
	// Initialize Utils
	utils.Initialize()

	// Initialize Routes
	log.Println("Initialize - Routes")
	initRoutes()

	// Initialize Server
	log.Println("Initialize - Server")
	mainServer = utils.NewServer(utils.RouterHandler)
}

// Main Function
func main() {
	// Starting Server
	log.Println("Server - Starting")
	mainServer.Start()

	// Make Channel to Catch OS Signal
	osSignal := make(chan os.Signal, 1)

	// Catch OS Signal from Channel
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)

	// Return OS Signal as Exit Code
	<-osSignal

	// Termination Symbol Log Line
	fmt.Println("")

	// Stopping Server
	log.Println("Server - Stopping")
	defer mainServer.Stop()

	// Close Any Database Connections
	if len(utils.Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(utils.Config.GetString("DB_DRIVER")) {
		case "mysql":
			defer utils.MySQL.Close()
		case "mongo":
			defer utils.MongoSession.Close()
		}
	}

	// Close Any Cache Connections
	if len(utils.Config.GetString("CACHE_DRIVER")) != 0 {
		switch strings.ToLower(utils.Config.GetString("CACHE_DRIVER")) {
		case "redis":
			defer utils.Redis.Close()
		}
	}
}
