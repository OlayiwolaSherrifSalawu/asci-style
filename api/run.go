package api

import (
	"log"
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
	
	a.configs.InfoLogger = log.New(os.Stdout, "INFO: \t", log.Ldate|log.Ltime)
	a.configs.ErrorLogger = log.New(os.Stderr, "ERROR: \t", log.Ldate|log.Ltime)
}
