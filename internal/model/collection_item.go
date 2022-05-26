package model

import (
	"time"
)

type CollectionItem struct {
	Id           int64     `db:"id,pk"`
	CollectionId int64     `db:"collection_id"`
	Value        string    `db:"value"`
	Index        int64     `db:"index"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
