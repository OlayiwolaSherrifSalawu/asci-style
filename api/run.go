package api

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func (a *Application) Run() {
	// setting up go routine
	// make your channel
	quit := make(chan os.Signal, 1)
	// create a signal to notify you when an issue happen
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	// creating flag
	a.InfoLogger = log.New(os.Stdout, "INFO: \t", log.Ldate|log.Ltime)
	a.ErrorLogger = log.New(os.Stderr, "ERROR: \t", log.Ldate|log.Ltime)
	// for changing port from the terminal with flag
	flag.StringVar(&a.Port, "port", ":5050", "Port Of server")
	flag.Parse()
	// create server and log it out
	mux := a.routes()
	server := http.Server{
		Addr:     a.Port,
		Handler:  mux,
		ErrorLog: a.ErrorLogger,
	}
	a.InfoLogger.Printf("started server at port %s \n", a.Port)
}
