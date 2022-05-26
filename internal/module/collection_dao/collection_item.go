package collection_dao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"aegis/internal/ecode"
	"aegis/internal/model"
	"aegis/pkg/mysqlx"
)

func isCollectionItemExists(ctx context.Context, db sqlx.ExtContext, collectionId int64, item string) (bool, error) {
	var count int64
	err := db.QueryRowxContext(
		ctx,
		`select count(*) from collection_items where collection_id = ? and value = ?`,
		collectionId,
		item,
	).Scan(&count)

	if err != nil {
		return true, err
	}
	return count != 0, nil
}

func CreateCollectionItem(ctx context.Context, db sqlx.ExtContext, collectionItem *model.CollectionItem) error {
	exists, err := isCollectionItemExists(ctx, db, collectionItem.CollectionId, collectionItem.Value)
	if err != nil {
		return fmt.Errorf("isCollectionItemExists: %w", err)
	}

	if exists {
		return ecode.ErrItemExists
	}

	err = mysqlx.Insert(ctx, db, "collection_items", collectionItem)
	if err != nil {
		return err
	}
	return err
}
