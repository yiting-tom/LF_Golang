package main

import (
	"11_http_multipart_requests/files"
	"11_http_multipart_requests/handlers"
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Log the debug level and above log messages.
	log.SetLevel(log.DebugLevel)
}

func main() {
	// Create a new logger
	l := log.New()

	// Set the base path for files storage
	sp, err := filepath.Abs("./local-storage/")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Create the local storage
	// max size is set to 5MB
	store, err := files.NewLocal(sp, 1024*1024*5)
	if err != nil {
		log.Fatal("failed to create local storage: ", err)
		os.Exit(1)
	}

	// Create the handler
	fh := handlers.NewFiles(l, store)

	// filename regex: {filename: [a-zA-Z]+\\.[a-z]{3}}
	fr := "/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{4}}"

	// Create a new serve mux and register the handler
	sm := mux.NewRouter()

	gh := sm.Methods(http.MethodGet).Subrouter()
	gh.HandleFunc(fr, fh.GetSingleImage)

	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc(fr, fh.Create)

	// Create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start the server
	go func() {
		l.Info("Start server on port 9090")

		if err := s.ListenAndServe(); err != nil {
			l.Fatal("Failed to start server: ", err)
			os.Exit(1)
		}
	}()

	// Trap the SIGINT signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// Block until we receive our signal.
	sig := <-c
	l.Info("Shutting down server with signal: ", sig)

	// Gracefully shutdown the server with a 30 second timeout
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
