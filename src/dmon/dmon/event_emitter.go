package dmon

import "code.cloudfoundry.org/lager"

type LoggingEventEmitter struct {
	Logger lager.Logger
}

func (e *LoggingEventEmitter) EmitEvent() error {
	e.Logger.Info("badness-occurred")
	return nil
}
