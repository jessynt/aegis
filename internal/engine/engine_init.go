package engine

import (
	"context"
)

func (e *Engine) Init(ctx context.Context) error {
	if err := e.store.Init(ctx); err != nil {
		_ = e.logger.Log("event", "store.init.failed")
		return err
	}
	_ = e.logger.Log("event", "store.init.done")
	if err := e.modelManager.Init(); err != nil {
		_ = e.logger.Log("event", "modelManager.init.failed")
		return err
	}
	_ = e.logger.Log("event", "modelManager.init.done")

	if err := e.propertyManager.Init(); err != nil {
		_ = e.logger.Log("event", "propertyManager.init.failed")
		return err
	}
	_ = e.logger.Log("event", "propertyManager.init.done")

	if err := e.abstractionManager.Init(); err != nil {
		_ = e.logger.Log("event", "abstractionManager.init.failed")
		return err
	}
	_ = e.logger.Log("event", "abstractionManager.init.done")

	if err := e.activationManager.Init(); err != nil {
		_ = e.logger.Log("event", "activationManager.init.failed")
		return err
	}
	_ = e.logger.Log("event", "activationManager.init.done")

	if err := e.ruleManager.Init(); err != nil {
		_ = e.logger.Log("event", "ruleManager.init.failed")
		return err
	}
	_ = e.logger.Log("event", "ruleManager.init.done")

	if err := e.collectionManager.Init(); err != nil {
		_ = e.logger.Log("event", "collectionManager.init.failed")
		return err
	}
	_ = e.logger.Log("event", "collectionManager.init.done")

	return nil
}
