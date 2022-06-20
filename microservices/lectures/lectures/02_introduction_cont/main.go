package main

import (
	"02_introduction_cont/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Binad address for the server")

// Official Documentation: https://golang.org/pkg/net/http/
func main() {
	// Create the Logger
	l := log.New(os.Stdout, "02_intro ", log.LstdFlags)

	// Create the handlers
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	// Create a new ServeMux and register handlers.
	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	// Create a new server with customized configurations.
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the server.
	go func() {
		l.Println("Starting server on port 9090")

		if err := server.ListenAndServe(); err != nil {
			l.Fatal(err)
			os.Exit(1)
		}
	}()

	// Trap sigterm or interupt and gracefully shotdown the server.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block unttil a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// Gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
