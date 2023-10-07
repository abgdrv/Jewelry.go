package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Application version_number
const version = "1.0.0"

// Configuration settings
type config struct {
	port int
	env  string
}

// Application struct to hold HTTP handlers, helpers, middleware
type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config

	// Parse configuration
	flag.IntVar(&cfg.port, "port", 4001, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialization of logger (recording information about the execution of an application)
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Declare instance of application
	app := application{config: cfg, logger: logger}

	// Declare HTTP server
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start HTTP server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
