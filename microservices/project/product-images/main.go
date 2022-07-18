package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/yiting-tom/LF_Golang/microservices/project/product-images/files"
	"github.com/yiting-tom/LF_Golang/microservices/project/product-images/handlers"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	l := logrus.New()

	store, err := files.NewLocal("./images/", 1024*1000)
	if err != nil {
		l.Fatal("error creating local store", "err", err)
		os.Exit(1)
	}

	// create a new Files handler
	handler := handlers.NewFiles(store, l)

	// create a server mux
	sm := mux.NewRouter()

	fileRegex := "/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}"

	// add the handler to the server mux
	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc(
		fileRegex,
		handler.Create,
	)
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc(
		fileRegex,
		handler.GetSingleImage,
	)

	// create a new server
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     &log.Logger{},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// start the server
	go func() {
		l.Info("Starting server on port 9090")

		if err := s.ListenAndServe(); err != nil {
			l.Fatal("error starting server", "err", err)
			os.Exit(1)
		}
	}()

	// trap SIGINT of SIGTERM signals to gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// block until we receive our signal
	sig := <-c
	l.Info("received signal", "signal", sig)

	// gracefully shutdown the server
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	s.Shutdown(ctx)
}
