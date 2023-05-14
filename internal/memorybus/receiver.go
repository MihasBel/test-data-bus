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
	msgList := mb.messages[message.Type]
	if len(msgList.Msgs) > mb.cfg.MaxMsgs {
		msgList.Msgs = msgList.Msgs[1:]
		msgList.ShiftCounter++
	} else {
		msgList.Msgs = append(msgList.Msgs, message)
	}
	return nil
}
