package rep

import (
	"context"

	"github.com/MihasBel/test-data-bus/internal/models"
)

type SubscriptionManager interface {
	Subscribe(ctx context.Context, subscriber *models.Subscriber) error
	Unsubscribe(ctx context.Context, id string, msgType string) error
}
