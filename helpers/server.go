package helpers

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Server Service Struct
type Server struct {
	srv *http.Server
	wg  sync.WaitGroup
}

// Server Configuration Struct
type ServerConfiguration struct {
	IP   string
	Port string
}

// Server Configuration Variable
var ServerConfig ServerConfiguration

// Function to Initialize New Server
func NewServer(handler http.Handler) *Server {
	// Initialize New Server
	return &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", ServerConfig.IP, ServerConfig.Port),
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
		fmt.Println("Service Started at", ServerConfig.IP+":"+ServerConfig.Port)
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
			fmt.Printf(fmt.Sprintf("Stopping Service Got an Error %v\n", err))
			return
		}
	}
	s.wg.Wait()
	fmt.Println("Service Stopped from", ServerConfig.IP+":"+ServerConfig.Port)
}
