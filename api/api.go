package api

import (
	"database/sql"
	"log"
)

func Run() {
	cons, err := ConnectDb("config.env")
	if err != nil {
		log.Fatal(err)
	}
	Db := sql.OpenDB(cons)
	app := NewApplication(Db)
	newApp := NewApplicationInterFace(app)
	mux := newApp.Routes()
	newApp.Run(mux)
}
