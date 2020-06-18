package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dimaskiddo/codebase-go-rest/hlp"
	"github.com/dimaskiddo/codebase-go-rest/hlp/cache"
	"github.com/dimaskiddo/codebase-go-rest/hlp/db"
	"github.com/dimaskiddo/codebase-go-rest/hlp/router"
)

// Server Variable
var svr *hlp.Server

// Init Function
func init() {
	// Initialize Server
	svr = hlp.NewServer(router.Router)
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
	if len(hlp.Config.GetString("DB_DRIVER")) != 0 {
		switch strings.ToLower(hlp.Config.GetString("DB_DRIVER")) {
		case "mysql":
			defer db.MySQL.Close()
		case "mongo":
			defer db.MongoSession.Close()
		}
	}

	// Close Any Cache Connections
	if len(hlp.Config.GetString("CACHE_DRIVER")) != 0 {
		switch strings.ToLower(hlp.Config.GetString("CACHE_DRIVER")) {
		case "redis":
			defer cache.Redis.Close()
		}
	}
}
