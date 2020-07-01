package controller

import (
	pb "gateway_grpc/protoc"
	"golang.org/x/net/context"
	"log"
	"time"
)

type EchoServer struct{}

func (s *EchoServer) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	log.Printf("EchoServer in %+v\n", in.String())

	return &pb.EchoReply{Message: "EchoServer Hello " + in.Name, Time: time.Now().String()}, nil
}
