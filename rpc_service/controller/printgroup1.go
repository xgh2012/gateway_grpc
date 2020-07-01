package controller

import (
	pb "gateway_grpc/protoc/group1"
	"golang.org/x/net/context"
	"log"
)

type PrintServerGroup1 struct{}

func (s *PrintServerGroup1) PrintGroup1(ctx context.Context, in *pb.PrintGroup1Request) (*pb.PrintGroup1Reply, error) {
	log.Printf("PrintServerGroup1 in %+v\n", in.String())

	return &pb.PrintGroup1Reply{Message: "PrintServerGroup1 Hello " + in.Name}, nil
}
