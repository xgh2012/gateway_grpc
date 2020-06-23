package main

import (
	"gateway_grpc/config"
	"gateway_grpc/gateway_service"
	"gateway_grpc/rpc_service"
	"time"
)

func main() {
	ch := make(chan bool, 1)
	go rpc_service.Run(ch)

	for {
		if config.RpcStatus == true && <-ch == true {
			gateway_service.Run()
		}
		time.Sleep(1 * time.Second)
	}
}
