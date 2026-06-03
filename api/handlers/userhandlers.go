package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"time"
)

type dto struct {
	UserId       string `json:"userid"`
	UserName     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"hashPassword"`
}
type UserRepo struct {
	Db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

func (a *UserRepo) CreateUser(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1234)
	contex, cancel := context.WithTimeout(r.Context(), time.Second*3)
	defer cancel()

}
