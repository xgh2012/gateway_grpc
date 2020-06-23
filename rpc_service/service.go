package rpc_service

import (
	"gateway_grpc/config"
	"log"
	"net"
	"strconv"
	"time"

	pb "gateway_grpc/gateway"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	log.Printf("in %+v\n", in.String())

	return &pb.EchoReply{Message: "Hello " + in.Name, Time: time.Now().String()}, nil
}

func Run(ch chan bool) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(int(config.GrpcPort)))

	if err != nil {
		log.Printf("failed to listen: %v", err)
		ch <- false
		return
	}

	s := grpc.NewServer()
	pb.RegisterGatewayServer(s, &server{})

	log.Println("Grpc 服务已经开启,监听端口：", config.GrpcPort)

	config.RpcStatus = true
	ch <- true

	err = s.Serve(lis)
	if err != nil {
		log.Printf("服务启动 failed: %v", err)
		ch <- false
		return
	}
}
