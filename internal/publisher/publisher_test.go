package publisher

import (
	"context"
	"testing"

	"github.com/MihasBel/test-data-bus/delivery/grpc/gen/v1/publisher"
	"github.com/MihasBel/test-data-bus/internal/models"
	"github.com/MihasBel/test-data-bus/mocks"
)

var (
	testSubscriber = &models.Subscriber{
		ID:          "testID",
		MessageType: "testType",
		Stream:      nil,
	}

	testMessage = &models.Message{
		Type: "testMsgType",
		Data: []byte("test data"),
	}
)

func TestService_Publish(t *testing.T) {
	s := New()
	_, cancel := context.WithCancel(context.Background())
	testSubscriber.Cancel = cancel
	streamSendCalled := false
	testSubscriber.Stream = &mocks.StreamMock{
		SendFunc: func(msg *publisher.Message) error {
			streamSendCalled = true
			return nil
		},
	}

	err := s.Publish(context.Background(), testSubscriber, testMessage)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !streamSendCalled {
		t.Errorf("Expected Stream.Send to be called, but it wasn't")
	}
}
