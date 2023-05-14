package subscription

import (
	"context"
	"fmt"
	"sync"

	"github.com/MihasBel/test-data-bus/internal/models"
	"github.com/MihasBel/test-data-bus/internal/rep"
)

const (
	subkeyf = "id_%s_type_%s"
)

type Service struct {
	subMap map[string]*models.Subscriber
	b      rep.Broker
	mu     sync.Mutex
}

func New(b rep.Broker) *Service {
	return &Service{
		b:      b,
		subMap: make(map[string]*models.Subscriber),
	}
}

func (s *Service) Subscribe(ctx context.Context, subscriber *models.Subscriber) error {
	err := subscriber.IsValid()
	if err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	key := fmt.Sprintf(subkeyf, subscriber.ID, subscriber.MessageType)
	v, ok := s.subMap[key]
	if ok {
		v.Cancel()
		v.Stream = subscriber.Stream
	} else {
		s.subMap[key] = subscriber
	}
	return s.b.HandleConsumer(ctx, s.subMap[key])
}

func (s *Service) Unsubscribe(_ context.Context, subscriberID string, msgType string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	key := fmt.Sprintf(subkeyf, subscriberID, msgType)
	v, ok := s.subMap[key]
	if ok {
		v.Cancel()
		return nil
	}
	return nil
}

// Start starts Service.
func (s *Service) Start(_ context.Context) error {
	return nil
}

// Stop stops Service.
func (s *Service) Stop(_ context.Context) error {
	return nil
}
