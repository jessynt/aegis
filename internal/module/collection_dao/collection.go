package collection_dao

import (
	"context"

	"github.com/jmoiron/sqlx"

	"aegis/internal/model"
	"aegis/pkg/mysqlx"
)

func CreateCollection(ctx context.Context, db sqlx.ExtContext, collection *model.Collection) error {
	err := mysqlx.Insert(ctx, db, "collections", collection)
	if err != nil {
		return err
	}
	return err
}
