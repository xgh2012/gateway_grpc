package route

import (
	pb "gateway_grpc/protoc"
	pg1 "gateway_grpc/protoc/group1"
	"gateway_grpc/rpc_service/controller"
	"google.golang.org/grpc"
	"log"
	"net"
)

//设置GRPC路由
func SetRouter(lis net.Listener) {
	//定义 grpc server
	server := grpc.NewServer()

	//设置 grpc server 路由
	Register(server)

	//启动 grpc server 服务
	err := server.Serve(lis)
	if err != nil {
		log.Printf("服务启动 PrintServer failed: %v", err)
		return
	}
}

//注册路由 方法
func Register(server *grpc.Server) {
	pb.RegisterGatewayPrintServer(server, &controller.PrintServer{})
	pb.RegisterGatewayEchoServer(server, &controller.EchoServer{})

	pg1.RegisterGatewayEchoGroup1Server(server, &controller.EchoServerGroup1{})
	pg1.RegisterGatewayPrintGroup1Server(server, &controller.PrintServerGroup1{})

}
