package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Server Service Struct
type Server struct {
	srv *http.Server
	wg  sync.WaitGroup
}

// Server Configuration Struct
type serverConfig struct {
	IP   string
	Port string
}

// Server Configuration Variable
var serverCfg serverConfig

// Function to Initialize New Server
func NewServer(handler http.Handler) *Server {
	// Initialize New Server
	return &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", serverCfg.IP, serverCfg.Port),
			Handler: handler,
		},
	}
}

// Method to Start Server
func (s *Server) Start() {
	// Initialize Context Handler Without Timeout
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Add to The WaitGroup for The Listener GoRoutine
	// And Wait for 1 Routine to be Done
	s.wg.Add(1)

	// Start The Server
	go func() {
		fmt.Println("Server Started at", serverCfg.IP+":"+serverCfg.Port)
		s.srv.ListenAndServe()

		s.wg.Done()
	}()
}

// Method to Stop Server
func (s *Server) Stop() {
	// Initialize Timeout
	timeout := 5 * time.Second

	// Initialize Context Handler With Timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Hanlde Any Error While Stopping Server
	if err := s.srv.Shutdown(ctx); err != nil {
		if err = s.srv.Close(); err != nil {
			log.Fatal("Error, " + strings.Title(err.Error()) + "!")
			return
		}
	}
	s.wg.Wait()
	fmt.Println("Server Stopped from", serverCfg.IP+":"+serverCfg.Port)
}
