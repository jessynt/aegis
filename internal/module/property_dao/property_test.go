package property_dao

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"

	"aegis/internal/dbtest"
	"aegis/internal/model"
)

func TestCreatePropertyMetadata(t *testing.T) {
	ctx := context.Background()

	t.Run("metadata create failed", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		require.NoError(t, err)
		sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
		defer mockDB.Close()
		f := dbtest.PropertyFactory.MustCreate().(*model.Property)

		mock.ExpectExec("INSERT INTO properties").WillReturnError(fmt.Errorf("error"))

		err = CreatePropertyMetadata(ctx, sqlxDB, f)
		require.Error(t, err)
	})
}
