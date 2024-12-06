package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config

	//Read value of the "port" and "env" command-line flags into config struct (4000 and development are default values)
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (dev | staging | prod)")

	// initialize logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// instance of app struct
	app := application{
		config: cfg,
		logger: logger,
	}

	// request dispactcher to paths

	// Declare HTTP server
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// start http server.
	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)

}