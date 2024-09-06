package adder

import (
	"context"

	"myGRPC/proto/createdProto"
)

type GRPCServer struct {
	createdProto.UnimplementedAdderServer
}

func (s *GRPCServer) SendMessage(ctx context.Context, in *createdProto.AddRequest) (*createdProto.AddResponse, error) {
	return &createdProto.AddResponse{Result: in.X + in.Y}, nil
}
