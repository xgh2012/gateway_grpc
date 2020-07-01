package rpc_service

import (
	"gateway_grpc/config"
	"gateway_grpc/rpc_service/route"
	"log"
	"net"
	"strconv"
)

func Run(ch chan bool) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(int(config.GrpcPort)))

	if err != nil {
		log.Printf("failed to listen: %v", err)
		ch <- false
		return
	}

	//设置路由
	go route.SetRouter(lis)

	log.Println("Grpc 服务已经开启,监听端口：", config.GrpcPort)

	config.RpcStatus = true
	ch <- true
}
