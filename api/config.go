package api

import (
	"database/sql"
	"log"
	"net/http"
)

type ApplicationInterface interface {
	Routes() *http.ServeMux
	Run(mux *http.ServeMux)
}
type Application struct {
	Db *sql.DB
	configs
}

type configs struct {
	Port        string
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

func NewApplication(db *sql.DB) *Application {
	return &Application{
		Db:      db,
		configs: configs{Port: ":4000", ErrorLogger: log.Default(), InfoLogger: log.Default()},
	}
}

func NewApplicationInterFace(app *Application) ApplicationInterface {
	return app
}
