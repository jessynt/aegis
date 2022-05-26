package model_dao

import (
	"context"

	"github.com/jmoiron/sqlx"

	"aegis/internal/model"
	"aegis/pkg/mysqlx"
)

// CreateModel 创建风控模型
func CreateModel(ctx context.Context, db sqlx.ExecerContext, model *model.Model) error {
	return mysqlx.Insert(ctx, db, "models", model)
}
