package ecode

import (
	"aegis/internal/model"
)

var (
	ErrInvalidArgument = model.MakeAppError(30000, "invalid argument")
	ErrItemExists      = model.MakeAppError(40000, "item already exists")
)
