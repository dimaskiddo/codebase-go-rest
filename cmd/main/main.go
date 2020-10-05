package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dimaskiddo/codebase-go-rest/pkg/cache"
	"github.com/dimaskiddo/codebase-go-rest/pkg/db"
	"github.com/dimaskiddo/codebase-go-rest/pkg/router"
	"github.com/dimaskiddo/codebase-go-rest/pkg/server"

	"github.com/dimaskiddo/codebase-go-rest/internal"
)

// Server Variable
var svr *server.Server

// Init Function
func init() {
	// Set Go Log Flags
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// Load Routes
	internal.LoadRoutes()

	// Initialize Server
	svr = server.NewServer(router.Router)
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
	if len(server.Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(server.Config.GetString("DB_DRIVER")) {
		case "mysql":
			defer db.MySQL.Close()
		case "mongo":
			defer db.MongoSession.Close()
		}
	}

	// Close Any Cache Connections
	if len(server.Config.GetString("CACHE_DRIVER")) != 0 {
		switch strings.ToLower(server.Config.GetString("CACHE_DRIVER")) {
		case "redis":
			defer cache.Redis.Close()
		}
	}
}
