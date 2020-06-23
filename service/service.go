package main

import (
	"log"
	"net"

	pb "gateway_grpc/gateway"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	PORT = ":9192"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("in %+v\n", in.String())

	return &pb.StringMessage{Value: "Hello " + in.Value}, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGatewayServer(s, &server{})
	log.Println("rpc服务已经开启")
	s.Serve(lis)
}
