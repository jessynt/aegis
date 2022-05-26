package validation

import (
	"fmt"
	"regexp"

	"github.com/spf13/cast"

	"aegis/internal/engine/contract"
	"aegis/internal/proto"
)

func Validate(v interface{}, p contract.Property) error {
	if IsReservedName(p.GetPropertyName()) {
		return fmt.Errorf("reserved key")
	}

	switch p.GetValidateType() {
	case proto.ValidateTypeNumber:
		if _, err := cast.ToFloat64E(v); err != nil {
			return fmt.Errorf("%#v 不是数字", v)
		}
	case proto.ValidateTypeString:
		if _, err := cast.ToStringE(v); err != nil {
			return fmt.Errorf("%#v 不是字符串", v)
		}
	case proto.ValidateTypeBool:
		if _, err := cast.ToBoolE(v); err != nil {
			return fmt.Errorf("%#v 不是布尔值", v)
		}
	case proto.ValidateTypeRegex:
		s, err := cast.ToStringE(v)
		if err != nil {
			return fmt.Errorf("%#v 无法识别", v)
		}
		matched, err := regexp.Match(p.GetValidateArgs(), []byte(s))
		if !matched || err != nil {
			return fmt.Errorf("%#v 不匹配", v)
		}
	case proto.ValidateTypeDatetime:
		s, err := cast.ToStringE(v)
		if err != nil {
			return fmt.Errorf("%#v 无法识别", v)
		}
		if _, err := cast.StringToDate(s); err != nil {
			return fmt.Errorf("%#v 不是时间", v)
		}
	default:
		return fmt.Errorf("不支持的验证类型")
	}
	return nil
}
