package memorybus

import (
	"context"
	"github.com/MihasBel/test-data-bus/internal/models"
	"github.com/pkg/errors"
	"time"
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
			if len(mb.messages[sub.MessageType]) == 0 ||
				sub.Offset < len(mb.messages[sub.MessageType]) {
				mb.log.Info().Msgf("No new messages for ID:%s", sub.ID)
				continue
			}
			err := mb.pub.Publish(ctx, sub, &models.Message{
				Type: sub.MessageType,
				Data: mb.messages[sub.MessageType][sub.Offset].Data,
			})
			if err != nil {
				mb.log.Error().Err(err).Msgf("Publish to subscriber ID:%v", sub.ID)
				continue
			}
			sub.Offset++
		}
	}
}
