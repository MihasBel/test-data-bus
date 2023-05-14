package rep

import (
	"context"

	"github.com/MihasBel/test-data-bus/internal/models"
)

type Receiver interface {
	ReceiveMsg(ctx context.Context, message models.Message) error
}
