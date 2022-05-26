package abstraction

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"aegis/internal/ecode"
	"aegis/internal/model"
	abstractionDao "aegis/internal/module/abstraction_dao"
)

type Service interface {
	CreateAbstraction(ctx context.Context, req createAbstractionRequest) error
}

type service struct {
	mysqlConn *sqlx.DB
}

func (s *service) CreateAbstraction(ctx context.Context, req createAbstractionRequest) error {
	if req.Name == "" || req.Label == "" {
		return ecode.ErrInvalidArgument
	}

	return abstractionDao.CreateAbstraction(ctx, s.mysqlConn, &model.Abstraction{
		Name:                   req.Name,
		Label:                  req.Label,
		ModelId:                req.ModelId,
		AggregateType:          int8(req.AggregateType),
		AggregateField:         req.AggregateField,
		AggregateIntervalType:  int8(req.AggregateIntervalType),
		AggregateIntervalValue: req.AggregateIntervalValue,
		SearchField:            req.SearchField,
		FilterExpression:       req.FilterExpression,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	})
}

func NewService(mysqlConn *sqlx.DB) *service {
	return &service{mysqlConn: mysqlConn}
}
