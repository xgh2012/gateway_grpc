package controller

import (
	pb "gateway_grpc/protoc"
	"golang.org/x/net/context"
	"log"
)

type PrintServer struct{}

func (s *PrintServer) Print(ctx context.Context, in *pb.PrintRequest) (*pb.PrintReply, error) {
	log.Printf("PrintServer in %+v\n", in.String())

	return &pb.PrintReply{Message: "PrintServer Hello " + in.Name}, nil
}
