package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/runtime/middleware"
	goHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	data "github.com/yiting-tom/LF_Golang/microservices/project/product-api/data"
	handlers "github.com/yiting-tom/LF_Golang/microservices/project/product-api/handlers"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	l := logrus.New()
	validation := data.NewValidation()

	handler := handlers.NewProducts(l, validation)

	sm := mux.NewRouter()

	getR := sm.Methods("GET").Subrouter()
	getR.HandleFunc("/products", handler.ListAll)
	getR.HandleFunc("/products/{id:[0-9]+}", handler.ListSingle)

	putR := sm.Methods("PUT").Subrouter()
	putR.HandleFunc("/products", handler.Update)
	putR.Use(handler.MiddlewareValidateProduct)

	postR := sm.Methods("POST").Subrouter()
	postR.HandleFunc("/products", handler.Create)
	putR.Use(handler.MiddlewareValidateProduct)

	deleteR := sm.Methods("DELETE").Subrouter()
	deleteR.HandleFunc("/products/{id:[0-9]+}", handler.Delete)

	// handler for swagger documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	ch := goHandlers.CORS(goHandlers.AllowedOrigins([]string{"*"}))

	// create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      ch(sm),
		ErrorLog:     log.New(l.Writer(), "", 0),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// start server
	go func() {
		l.Info("Starting server on port :9090")
		if err := s.ListenAndServe(); err != nil {
			l.Fatal(err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Info("Got signal: ", sig)

	// gracefully shutdown the server, waiting max. 10 seconds for current operations to complete
	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		l.Fatal(err)
	}
	s.Shutdown(ctx)
}
