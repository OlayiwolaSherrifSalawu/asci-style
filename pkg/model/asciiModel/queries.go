package asciimodel

import (
	"context"
	"database/sql"

	"github.com/OlayiwolaSherrifSalawu/asci-style.git/pkg/model"
)

type AsciModel struct {
	Db *sql.DB
}

func (d *AsciModel) Insert(ctx context.Context, Users *model.User) error {

}
