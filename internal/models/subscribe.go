package models

import (
	"context"

	"github.com/MihasBel/test-data-bus/delivery/grpc/gen/v1/publisher"
	"github.com/pkg/errors"
)

var (
	ErrEmptID = errors.New("subscriber ID cannot be empty")
)

type Subscriber struct {
	ID           string
	Stream       publisher.PubSubService_SubscribeServer
	MessageType  string
	Cancel       context.CancelFunc
	Offset       int
	ShiftCounter int
}

func (s *Subscriber) IsValid() error {
	if s.ID == "" {
		return ErrEmptID
	}
	if s.Stream == nil {
		return errors.New("stream cannot be nil")
	}
	if s.MessageType == "" {
		return errors.New("message type cannot be empty")
	}
	if s.Cancel == nil {
		return errors.New("CancelFunc ch cannot be nil")
	}
	return nil
}
