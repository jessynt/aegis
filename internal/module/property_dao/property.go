package property_dao

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"aegis/internal/model"
	"aegis/pkg/mysqlx"
)

var (
	_alterEventTableSQL = "ALTER TABLE event_dist ADD COLUMN %s %s"
)

// CreatePropertyMetadata 创建属性元数据
func CreatePropertyMetadata(ctx context.Context, db sqlx.ExecerContext, property *model.Property) error {
	// 1. Update metadata
	err := mysqlx.Insert(ctx, db, "properties", property)
	if err != nil {
		return err
	}
	return err
}

// 创建数据仓库字段
func CreateWarehouseField(ctx context.Context, ch sqlx.ExecerContext, name, dataType string) error {
	_, err := ch.ExecContext(ctx, fmt.Sprintf(_alterEventTableSQL, name, dataType))
	return err
}
