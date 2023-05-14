package rep

import (
	"context"

	"github.com/MihasBel/test-data-bus/internal/models"
)

type Publisher interface {
	Publish(ctx context.Context, sub *models.Subscriber, msg *models.Message) error
}
