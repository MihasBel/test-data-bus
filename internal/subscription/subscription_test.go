package subscription

import (
	"context"
	"testing"

	"github.com/MihasBel/test-data-bus/delivery/grpc/gen/v1/publisher"
	"github.com/MihasBel/test-data-bus/internal/models"
	"github.com/MihasBel/test-data-bus/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
)

func TestService_Subscribe(t *testing.T) {
	stream := &mocks.StreamMock{
		SendFunc: func(msg *publisher.Message) error {
			return nil
		},
	}
	_, cancel := context.WithCancel(context.Background())
	tests := []struct {
		name        string
		subscriber  *models.Subscriber
		wantErr     error
		brokerError error
	}{
		{
			name: "valid subscriber, not yet subscribed",
			subscriber: &models.Subscriber{
				ID:          "123",
				MessageType: "type1",
				Stream:      stream,
				Cancel:      cancel,
			},
			wantErr: nil,
		},
		{
			name: "valid subscriber, already subscribed",
			subscriber: &models.Subscriber{
				ID:          "123",
				MessageType: "type1",
				Stream:      stream,
				Cancel:      cancel,
			},
			wantErr: nil,
		},
		{
			name: "invalid subscriber",
			subscriber: &models.Subscriber{
				ID:          "",
				MessageType: "type1",
				Stream:      stream,
				Cancel:      cancel,
			},
			wantErr: models.ErrEmptID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			broker := mocks.BrokerMock{}
			if tt.wantErr == nil {
				broker.On("HandleConsumer", mock.Anything, tt.subscriber).Return(tt.brokerError)
			}
			service := New(&broker)

			err := service.Subscribe(context.Background(), tt.subscriber)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Service.Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}

			broker.AssertExpectations(t)
		})
	}
}
