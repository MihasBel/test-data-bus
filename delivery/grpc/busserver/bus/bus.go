package bus

import (
	"github.com/MihasBel/test-data-bus/delivery/grpc/gen/v1/bus"
	"github.com/MihasBel/test-data-bus/internal/rep"
)

// Server receiver
type Server struct {
	bus.UnimplementedBusServiceServer
	r rep.Receiver
}

// New constructor
func New(r rep.Receiver) *Server {
	return &Server{
		r: r,
	}
}
