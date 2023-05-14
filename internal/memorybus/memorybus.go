package memorybus

import (
	"context"
	"sync"

	"github.com/MihasBel/test-data-bus/internal/models"
	"github.com/MihasBel/test-data-bus/internal/rep"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const (
	MsgErrUnknownType = "unknown message type in received message: %v"
	MsgErrNoTypes     = "No message types are specified for processing"
	MsgErrEmptyMsg    = "empty message data in received message type: %v"
)

type MemoryBus struct {
	cfg      Config
	log      *zerolog.Logger
	messages map[string][]models.Message
	pub      rep.Publisher
	mu       sync.Mutex
}

func New(cfg Config, log *zerolog.Logger, pub rep.Publisher) *MemoryBus {
	return &MemoryBus{
		cfg:      cfg,
		log:      log,
		messages: make(map[string][]models.Message),
		pub:      pub,
	}
}

// Start MemoryBus
func (mb *MemoryBus) Start(_ context.Context) error {
	if len(mb.cfg.MsgTypes) == 0 {
		return errors.New(MsgErrNoTypes)
	}
	for _, v := range mb.cfg.MsgTypes {
		mb.messages[v] = make([]models.Message, 0)
	}
	return nil
}

// Stop MemoryBus
func (mb *MemoryBus) Stop(_ context.Context) error {
	return nil
}

func contains(slice []string, str string) bool {
	for _, a := range slice {
		if a == str {
			return true
		}
	}
	return false
}
