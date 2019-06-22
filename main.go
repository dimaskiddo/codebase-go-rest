package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	svc "github.com/dimaskiddo/codebase-go-rest/service"
)

// Server Variable
var svr *svc.Server

// Init Function
func init() {
	// Initialize Routes
	routesInit()

	// Initialize Server
	svr = svc.NewServer(svc.Router)
}

// Main Function
func main() {
	// Starting Server
	svr.Start()

	// Make Channel for OS Signal
	sig := make(chan os.Signal, 1)

	// Notify Any Signal to OS Signal Channel
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	// Return OS Signal Channel
	// As Exit Sign
	<-sig

	// Log Break Line
	fmt.Println("")

	// Stopping Server
	defer svr.Stop()

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
