package memorybus

import (
	"context"
	"github.com/MihasBel/test-data-bus/internal/models"
	"github.com/pkg/errors"
)

func (mb *MemoryBus) ReceiveMsg(_ context.Context, message models.Message) error {
	if !contains(mb.cfg.MsgTypes, message.Type) {
		return errors.Errorf(MsgErrUnknownType, message.Type)
	}
	if len(message.Data) == 0 {
		return errors.Errorf(MsgErrEmptyMsg, message.Type)
	}
	mb.mu.Lock()
	defer mb.mu.Unlock()
	mb.messages[message.Type] = append(mb.messages[message.Type], message)
	return nil
}
