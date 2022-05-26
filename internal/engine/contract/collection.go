package contract

import (
	"aegis/internal/proto"
)

type Collection interface {
	Name() string
	Contains(k interface{}) bool
	GetPropertyId() int64
	Type() proto.CollectionType
}
