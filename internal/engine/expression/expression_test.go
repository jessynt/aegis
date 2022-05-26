package expression

import (
	"errors"
	"fmt"
	"testing"

	"github.com/antonmedv/expr"
	"github.com/stretchr/testify/require"
)

func TestCheck(t *testing.T) {
	type arg struct {
		expression string
		env        map[string]interface{}
	}
	tests := []struct {
		name       string
		arg        arg
		wantResult bool
		wantError  error
	}{
		{
			name: "invalid expression 1",
			arg: arg{
				expression: "1 + 2",
				env:        nil,
			},
			wantResult: false,
			wantError:  ErrInvalidExpression,
		},
		{
			name: "invalid expression 2",
			arg: arg{
				expression: "1 + foo == 3",
				env: map[string]interface{}{
					"foo": 1,
				},
			},
			wantResult: false,
			wantError:  nil,
		},
		{
			name: "valid expression 1",
			arg: arg{
				expression: "1 < 3",
				env:        nil,
			},
			wantResult: true,
			wantError:  nil,
		},
		{
			name: "valid expression 2",
			arg: arg{
				expression: "1 < 3 and 3 > 2",
				env:        nil,
			},
			wantResult: true,
			wantError:  nil,
		},
		{
			name: "valid expression 3",
			arg: arg{
				expression: "1 * login_ip_10_min_qty < 100",
				env: map[string]interface{}{
					"login_ip_10_min_qty": 20,
				},
			},
			wantResult: true,
			wantError:  nil,
		},
		{
			name: "valid expression 4",
			arg: arg{
				expression: "3 * login_ip_10_min_qty < 10",
				env: map[string]interface{}{
					"login_ip_10_min_qty": 5,
				},
			},
			wantResult: false,
			wantError:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rv, err := Check(tt.arg.expression, tt.arg.env)
			require.True(t, errors.Is(err, tt.wantError))
			require.Equal(t, tt.wantResult, rv)
		})
	}
}

func Test2(t *testing.T) {
	rv, err := expr.Eval(`1 in ["1"]`, map[string]interface{}{})
	require.NoError(t, err)
	fmt.Printf("%+v\n", rv)
}
