package expression

import (
	"fmt"

	"github.com/antonmedv/expr"
	"github.com/pkg/errors"
)

var (
	ErrInvalidExpression = errors.New("invalid expression")
)

// Check 检查数据是否满足表达式要求
func Check(expression string, env interface{}) (rv bool, err error) {
	result, err := expr.Eval(expression, env)
	if err != nil {
		return false, err
	}
	rv, ok := result.(bool)
	if !ok {
		return false, fmt.Errorf("%w: %s", ErrInvalidExpression, expression)
	}
	return rv, nil
}
