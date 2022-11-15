package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/opensaucerer/gotemplate/config"
	"github.com/opensaucerer/gotemplate/version"

	service "github.com/opensaucerer/gotemplate/service"

	handlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func createServer() (s *http.Server) {
	r := mux.NewRouter()
	// inject combined logger (apache & nginx style)
	logger := handlers.CombinedLoggingHandler(os.Stdout, r)

	// we should do more cross origin stuff here
	r.Use(handlers.CORS())

	// register routes with versioning
	version.VersionRoutes(r)

	port := fmt.Sprintf(":%s", config.MustGet("PORT", "8080"))

	s = &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        logger,
	}

	go func() {
		service.MustInitiateDynamo()
		log.Printf("Starting at http://127.0.0.1%s", port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error listening on port: %s\n", err)
		}
	}()

	return s
}

func main() {

	s := createServer()
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shut down...")
	}
	log.Println("Server exited!")
}
