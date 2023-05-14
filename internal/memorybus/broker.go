package memorybus

import (
	"context"
	"time"

	"github.com/MihasBel/test-data-bus/internal/models"
	"github.com/pkg/errors"
)

func (mb *MemoryBus) HandleConsumer(ctx context.Context, subscriber *models.Subscriber) error {
	if !contains(mb.cfg.MsgTypes, subscriber.MessageType) {
		return errors.Errorf(MsgErrUnknownType, subscriber.MessageType)
	}
	go mb.consume(ctx, subscriber)

	return nil
}

func (mb *MemoryBus) consume(ctx context.Context, sub *models.Subscriber) {
	for range time.NewTicker(mb.cfg.ReadDelaySec).C {
		select {
		case <-ctx.Done():
			mb.log.Info().Msgf("cancel ctx in HandleConsumer go func id:%s", sub.ID)
			return
		default:
			mb.mu.Lock()
			msgList := mb.messages[sub.MessageType]
			if len(msgList.Msgs) == 0 {
				mb.mu.Unlock()
				continue
			}
			missedMsgs := msgList.ShiftCounter - sub.ShiftCounter
			if missedMsgs > 0 {
				sub.Offset -= missedMsgs
				sub.ShiftCounter = msgList.ShiftCounter
			}
			if sub.Offset >= len(msgList.Msgs) {
				mb.log.Info().Msgf("No new messages for ID:%s", sub.ID)
				mb.mu.Unlock()
				continue
			}
			if sub.Offset < 0 {
				sub.Offset = 0
			}
			msg := msgList.Msgs[sub.Offset]
			mb.mu.Unlock()
			err := mb.pub.Publish(ctx, sub, &models.Message{
				Type: msg.Type,
				Data: msg.Data,
			})
			if err != nil {
				mb.log.Error().Err(err).Msgf("Publish to subscriber ID:%v", sub.ID)
				continue
			}
			sub.Offset++
		}
	}
}
