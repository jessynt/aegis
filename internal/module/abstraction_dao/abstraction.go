package abstraction_dao

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"

	"aegis/internal/model"
	"aegis/pkg/mysqlx"
)

var (
	ErrNameExists = errors.New("abstraction name already exists")
)

func isAbstractionNameExists(ctx context.Context, db sqlx.ExtContext, name string) (bool, error) {
	var count int64
	err := db.QueryRowxContext(
		ctx,
		`select count(*) from abstractions where name = ?`,
		name,
	).Scan(&count)

	if err != nil {
		return true, err
	}
	return count != 0, nil
}

func CreateAbstraction(ctx context.Context, db sqlx.ExtContext, abstraction *model.Abstraction) error {
	exists, err := isAbstractionNameExists(ctx, db, abstraction.Name)
	if err != nil {
		return err
	}
	if exists {
		return ErrNameExists
	}
	return mysqlx.Insert(ctx, db, "abstractions", abstraction)
}
