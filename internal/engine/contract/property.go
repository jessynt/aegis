package contract

import (
	"aegis/internal/proto"
)

type Property interface {
	Id() int64
	GetType() proto.PropertyType
	GetValidateType() proto.ValidateType
	GetValidateArgs() string
	GetPropertyName() string
}
