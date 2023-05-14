package rep

import (
	"context"

	"github.com/MihasBel/test-data-bus/internal/models"
)

type Broker interface {
	HandleConsumer(ctx context.Context, subscriber *models.Subscriber) error
}
