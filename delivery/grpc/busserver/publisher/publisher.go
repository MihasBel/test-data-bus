package publisher

import (
	"github.com/MihasBel/test-data-bus/delivery/grpc/gen/v1/publisher"
	"github.com/MihasBel/test-data-bus/internal/rep"
)

// Server SubscriptionManager
type Server struct {
	publisher.UnimplementedPubSubServiceServer
	m rep.SubscriptionManager
}

// New constructor
func New(m rep.SubscriptionManager) *Server {
	return &Server{
		m: m,
	}
}
