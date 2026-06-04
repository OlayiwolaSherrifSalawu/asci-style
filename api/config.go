package api

import (
	"log"

	"github.com/OlayiwolaSherrifSalawu/asci-style.git/api/handlers"
)

type ApplicationInterface interface {
	Run()
}
type Application struct {
	AsciiHandler *handlers.AsciHandlers
	UserHandler  *handlers.UserRepo
	configs
}

type configs struct {
	Port        string
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func NewApplication(port string, userRepor *handlers.UserRepo, asciiHandler *handlers.AsciHandlers) *Application {
	return &Application{
		UserHandler:  userRepor,
		configs:      configs{Port: port, ErrorLogger: log.Default(), InfoLogger: log.Default()},
		AsciiHandler: asciiHandler,
	}
}

func NewApplicationInterFace(app *Application) ApplicationInterface {
	return app
}
