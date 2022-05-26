package model

import (
	"time"
)

type Collection struct {
	Id         int64     `db:"id,pk"`
	Name       string    `db:"name"`
	Label      string    `db:"label"`
	PropertyId int64     `db:"property_id"`
	Comment    string    `db:"comment"`
	Type       int32     `db:"type"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
