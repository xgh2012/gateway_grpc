package controller

import (
	pb "gateway_grpc/protoc/group1"
	"golang.org/x/net/context"
	"log"
	"time"
)

type EchoServerGroup1 struct{}

func (s *EchoServerGroup1) EchoGroup1(ctx context.Context, in *pb.EchoGroup1Request) (*pb.EchoGroup1Reply, error) {
	log.Printf("EchoServerGroup1 in %+v\n", in.String())

	return &pb.EchoGroup1Reply{Message: "EchoServerGroup1 Hello " + in.Name, Time: time.Now().String()}, nil
}
