package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Server configuration struct
type config struct {
	port int
	env  string
}

// Application dependencies
type application struct {
	cfg    config
	logger *log.Logger
}

func main() {
	// Define flags
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "server network port")
	flag.StringVar(&cfg.env, "environment", "development", "server environment")
	flag.Parse()

	// Define loggers
	logger := log.New(os.Stdout, "INFO: ", log.Ltime)

	app := &application{
		cfg:    cfg,
		logger: logger,
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	// Starting Server
	logger.Printf("starting server on port %s", srv.Addr)
	logger.Fatal(srv.ListenAndServe())
}
