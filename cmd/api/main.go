package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// Declare application version number , currently hard coded but will generate this
// at build time.
const version = "1.0.0"

// Holds configuration data needed to run application
type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {

	// Initialize a new instance of config
	var cfg config

	// Parse cli arguments and set values of config if nothing is set port is defaulted to 8080
	// environment is set to development
	flag.IntVar(&cfg.port, "port", 8080, "port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare server with mux we configured and set sensible timeouts
	// also includes error logging.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	// Server listens to port that has been set
	err := srv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
	// Exit if fail
	os.Exit(1)

}
