package collection

import (
	"github.com/spf13/cast"

	"aegis/internal/model"
	"aegis/internal/proto"
	"aegis/pkg/hashset"
)

type Collection struct {
	c   *model.Collection
	set *hashset.Set // Should we using LinkedHashSet?
}

func NewFromModel(dbCollection *model.Collection, items []*model.CollectionItem) *Collection {
	c := &Collection{
		c:   dbCollection,
		set: hashset.New(),
	}
	for _, item := range items {
		c.set.Add(item.Value)
	}
	return c
}

func (c Collection) Name() string {
	return c.c.Name
}

func (c *Collection) GetPropertyId() int64 {
	return c.c.PropertyId
}

func (c *Collection) Contains(value interface{}) bool {
	return c.set.Contains(cast.ToString(value)) // FIXME: value 无法转换成 string 时可能被绕过
}

func (c *Collection) Type() proto.CollectionType {
	return proto.CollectionType(c.c.Type)
}
