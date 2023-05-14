package publisher

import (
	"context"

	"github.com/MihasBel/test-data-bus/delivery/grpc/gen/v1/publisher"
	"github.com/MihasBel/test-data-bus/internal/models"
	"github.com/pkg/errors"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Publish(_ context.Context, subscriber *models.Subscriber, msg *models.Message) error {
	err := subscriber.IsValid()
	if err != nil {
		return err
	}
	grpcMsg := &publisher.Message{
		Type: msg.Type,
		Data: msg.Data,
	}

	if err := subscriber.Stream.Send(grpcMsg); err != nil {
		return errors.Wrapf(err, "failed to send message to subscriber %s", subscriber.ID)
	}

	return nil
}
