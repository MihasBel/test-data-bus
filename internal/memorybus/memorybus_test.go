package memorybus

import (
	"context"
	"testing"
	"time"

	"github.com/MihasBel/test-data-bus/internal/models"
	"github.com/MihasBel/test-data-bus/mocks"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestHandleConsumer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPublisher := mocks.NewMockPublisher(ctrl)

	msgTypes := []string{"type1"}
	cfg := Config{MsgTypes: msgTypes, MaxMsgs: 5, ReadDelayMS: time.Second}
	logger := zerolog.Nop()

	mb := New(cfg, &logger, mockPublisher)

	// Starting memory bus with the message types
	err := mb.Start(context.Background())
	assert.Nil(t, err)

	sub := &models.Subscriber{ID: "sub1", MessageType: "type1", Offset: 2, ShiftCounter: 0}
	mockPublisher.EXPECT().Publish(gomock.Any(), sub, gomock.Any()).Return(nil).AnyTimes()
	// Simulating the receipt of messages
	for i := 0; i < 10; i++ {
		mb.ReceiveMsg(context.Background(), models.Message{Type: "type1", Data: []byte("message")})
	}
	mb.HandleConsumer(context.Background(), sub)

	time.Sleep(time.Second * 5)

	assert.Equal(t, 4, sub.Offset)
	assert.Equal(t, 5, sub.ShiftCounter)
}
