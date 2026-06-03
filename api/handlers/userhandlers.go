package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/OlayiwolaSherrifSalawu/asci-style.git/pkg/model"
	asciimodel "github.com/OlayiwolaSherrifSalawu/asci-style.git/pkg/model/asciiModel"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type dto struct {
	UserId       string `json:"userid"`
	UserName     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"hashPassword"`
}
type UserRepo struct {
	Db     *sql.DB
	UModel *asciimodel.UserModel
}

func NewUserHandler(db *sql.DB) *UserRepo {
	return &UserRepo{
		Db:     db,
		UModel: &asciimodel.UserModel{Db: db},
	}
}

func (a *UserRepo) CreateUser(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1234)
	dtos := &dto{}
	err := json.NewDecoder(r.Body).Decode(dtos)
	if err != nil {
		return
	}
	dtos.UserId = uuid.NewString()
	contex, cancel := context.WithTimeout(r.Context(), time.Second*3)
	defer cancel()
	pass, err := bcrypt.GenerateFromPassword([]byte(dtos.Password), 12)
	dtos.Password = string(pass)
	users := &model.User{
		UserId:       dtos.UserId,
		UserName:     dtos.UserName,
		EmailAddress: dtos.EmailAddress,
		HashPassword: dtos.Password,
	}
	dtos = nil
	err = a.UModel.Insert(contex, users)
	if err != nil {
		return
	}
	users = nil
	fmt.Fprint(w, "created users")
}
