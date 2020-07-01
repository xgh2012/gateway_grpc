package client

import (
	pb "gateway_grpc/protoc"
	pg1 "gateway_grpc/protoc/group1"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//gateway http 服务设置
func SetClientRouter(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	var err error
	err = pb.RegisterGatewayEchoHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	err = pb.RegisterGatewayPrintHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	err = pg1.RegisterGatewayEchoGroup1HandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	err = pg1.RegisterGatewayPrintGroup1HandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	return nil
}
