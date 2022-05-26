package model

import (
	"context"
	"time"

	gkitLog "github.com/go-kit/kit/log"
)

type loggingService struct {
	logger gkitLog.Logger
	Service
}

func NewLoggingService(logger gkitLog.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) CreateModel(ctx context.Context, guid, name, label string) (err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			"method", "create",
			"name", name,
			"GUID", guid,
			"label", label,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.CreateModel(ctx, guid, name, label)

}
