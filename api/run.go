package api

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	flag.StringVar(&a.Port, "port", a.Port, "Port Of server")
	flag.Parse()
	// create server and log it out
	mux := a.routes()
	server := http.Server{
		Addr:     a.Port,
		Handler:  mux,
		ErrorLog: a.ErrorLogger,
	}
	go func() {
		a.InfoLogger.Printf("started server at port %s \n", a.Port)
		err := server.ListenAndServe()
		a.InfoLogger.Printf("started server at port %s \n", a.Port)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.ErrorLogger.Println(err)
			return
		}
	}()
	<-quit
	a.InfoLogger.Println("Shutting Down server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := server.Shutdown(ctx)

	if err != nil {
		a.ErrorLogger.Println(err)
		return
	}
	a.InfoLogger.Println("server shutdown properly")
}
