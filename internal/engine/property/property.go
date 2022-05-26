package property

import (
	"aegis/internal/model"
	"aegis/internal/proto"
)

type Property struct {
	m *model.Property
}

func NewPropertyFromModel(m *model.Property) *Property {
	return &Property{
		m: m,
	}
}

// Id 获取属性 ID
func (f Property) Id() int64 {
	return f.m.Id
}

// GetType 获取属性数据类型
func (f Property) GetType() proto.PropertyType {
	return proto.PropertyType(f.m.Type)
}

// GetValidateType 获取属性验证类型
func (f Property) GetValidateType() proto.ValidateType {
	return proto.ValidateType(f.m.ValidateType)
}

// GetValidateArgs 获取属性验证参数
func (f Property) GetValidateArgs() string {
	return f.m.ValidateArgs
}

// GetValidateArgs 获取属性名称
func (f Property) GetPropertyName() string {
	return f.m.Name
}

func (f Property) String() string {
	return f.m.Name
}
