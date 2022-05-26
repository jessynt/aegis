package event_dao

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

func InsertEvent(ctx context.Context, clickhouseConn *sqlx.DB, record map[string]interface{}) error {
	var names []string
	var values []interface{}
	for name, value := range record {
		names = append(names, quoted(name))
		values = append(values, value)
	}

	insertQuery := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		quoted("event_dist"),
		strings.Join(names, ","),
		strings.Repeat(placeholderString+",", len(names)-1)+placeholderString,
	)

	tx, err := clickhouseConn.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, insertQuery, values...)
	if err != nil {
		return err
	}
	return tx.Commit()
}
