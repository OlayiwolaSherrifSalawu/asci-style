package handlers

import (
	"database/sql"
	"net/http"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

func (a *UserRepo) CreateUser(w http.ResponseWriter, r *http.Request){
	
}
