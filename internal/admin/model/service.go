package model

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"aegis/internal/ecode"
	"aegis/internal/model"
	modelDao "aegis/internal/module/model_dao"
)

type Service interface {
	CreateModel(ctx context.Context, guid, name, label string) error
}

type service struct {
	mysqlConn *sqlx.DB
}

func (s *service) CreateModel(ctx context.Context, guid, name, label string) error {
	if guid == "" || name == "" || label == "" {
		return ecode.ErrInvalidArgument
	}

	return modelDao.CreateModel(ctx, s.mysqlConn, &model.Model{
		Name:      name,
		Label:     label,
		GUID:      guid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func NewService(mysqlConn *sqlx.DB) Service {
	return &service{mysqlConn: mysqlConn}
}
