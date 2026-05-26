package handlers

import "database/sql"

type UserRepo struct {
	Db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}
