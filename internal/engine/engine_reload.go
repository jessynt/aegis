package engine

import (
	"context"
)

func (e *Engine) Reload(ctx context.Context) error {
	return e.Init(ctx)
}
