package model_dao

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"

	"aegis/internal/model"
)

func TestCreateModel(t *testing.T) {
	ctx := context.Background()

	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	m := &model.Model{
		Name:      "login",
		Label:     "登录行为",
		GUID:      uuid.New().String(),
		Status:    0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	mock.
		ExpectExec("INSERT INTO `models` (`name`,`label`,`guid`,`status`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?)").
		WithArgs(m.Name, m.Label, m.GUID, m.Status, m.CreatedAt, m.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = CreateModel(
		ctx,
		sqlxDB, m,
	)
	require.NoError(t, err)
}
