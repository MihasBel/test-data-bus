package bus

import (
	"context"

	"github.com/MihasBel/test-data-bus/delivery/grpc/gen/v1/bus"
	"github.com/MihasBel/test-data-bus/internal/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ReceiveMessage(ctx context.Context, msg *bus.Message) (*emptypb.Empty, error) {
	if err := s.r.ReceiveMsg(ctx, models.Message{
		Type: msg.Type,
		Data: msg.Data,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to receive message: %v", err)
	}
	return &emptypb.Empty{}, nil
}
