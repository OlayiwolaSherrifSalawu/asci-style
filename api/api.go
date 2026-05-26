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

}
