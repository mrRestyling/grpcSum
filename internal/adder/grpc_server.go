package adder

import (
	"context"

	"myGRPC/proto/createdProto"
)

// GRPCServer - реализует интерфейс AdderServer из пакета createdProto
type GRPCServer struct {
	createdProto.UnimplementedAdderServer
}

// SendMessage(метод) является обработчиком функции для метода RPC SendMessage,
// определенного в сервисе Adder в файле протокола.
func (s *GRPCServer) SendMessage(ctx context.Context, in *createdProto.AddRequest) (*createdProto.AddResponse, error) {
	return &createdProto.AddResponse{Result: in.X + in.Y}, nil
}
