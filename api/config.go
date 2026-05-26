package api

import (
	"log"
	"net/http"

	"github.com/OlayiwolaSherrifSalawu/asci-style.git/api/handlers"
)

type ApplicationInterface interface {
	Routes() *http.ServeMux
	Run(mux *http.ServeMux)
}
type Application struct {
	UserHandler *handlers.UserRepo
	configs
}

type configs struct {
	Port        string
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func NewApplication(userRepor *handlers.UserRepo) *Application {
	return &Application{
		UserHandler: userRepor,
		configs:     configs{Port: ":4000", ErrorLogger: log.Default(), InfoLogger: log.Default()},
	}
}

func NewApplicationInterFace(app *Application) ApplicationInterface {
	return app
}
