package asciimodel

import (
	"context"
	"database/sql"

	"github.com/OlayiwolaSherrifSalawu/asci-style.git/pkg/model"
)

type UserModel struct {
	Db *sql.DB
}

func (d *UserModel) Insert(ctx context.Context, Users *model.User) error {
	stmt := "INSERT INTO users(id, user_name, email_address, password) VALUES($1, $2, $3, $4)"
	tx, err := d.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, stmt, Users.UserId, Users.UserName, Users.HashPassword)
	if err != nil {
		return err
	}
	return nil
}
