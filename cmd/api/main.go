package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// Server configuration struct
type config struct {
	port int
	env  string

	db struct {
		dsn string
	}
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
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("WRITEM_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	// Define loggers
	logger := log.New(os.Stdout, "INFO: ", log.Ltime)

	// Create a connection pool
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

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

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
