package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dimaskiddo/frame-go/dbs"
	"github.com/dimaskiddo/frame-go/utils"
)

func main() {
	// Initialize Channel for OS Signal
	signalOS := make(chan os.Signal, 1)

	// Initialize Configuration
	utils.InitConfig()

	// Initialize Database
	if len(utils.Config.GetString("DB_DRIVER")) != 0 {
		utils.InitDB()
	}

	// Initialize Router
	utils.InitRouter()

	// Initialize Routes
	InitRoutes()

	// Initialize Server
	server := utils.NewServer(utils.RouterHandler)

	// Starting Server
	server.Start()

	// Catch OS Signal from Channel
	signal.Notify(signalOS, os.Interrupt, syscall.SIGTERM)

	// Return OS Signal as Exit Code
	<-signalOS

	// Add Some Spaces When Done
	fmt.Println(" Stopping Server ")

	// Defer Some Function Before End
	defer server.Stop()
	switch strings.ToLower(utils.Config.GetString("DB_DRIVER")) {
	case "mysql":
		defer dbs.MySQL.Close()
	case "mongo":
		defer dbs.MongoSession.Close()
	}
}
