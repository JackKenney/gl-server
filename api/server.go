package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"time"
)

// Server defines the router and the configuration variables.
type Server struct {
	router Router
	wait   time.Duration
}

// ProvideServer returns the location of the server instance on the heap.
func ProvideServer(router Router, wait time.Duration) *Server {
	return &Server{
		router: router,
		wait:   wait,
	}
}

// Run starts the server listening with the default parameters
func (s *Server) Run(port string) {
	srv := &http.Server{
		Addr:    port,
		Handler: s.router, // Pass our instance of gorilla/mux in.
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Listening on port", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	// Create a deadline to wait for.
	wait := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
