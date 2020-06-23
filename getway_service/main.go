package main

import (
	"flag"
	"gateway_grpc/config"
	"log"
	"net/http"
	"strconv"

	gw "gateway_grpc/gateway"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:"+strconv.Itoa(int(config.ServicePort)), "endpoint of Gateway")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterGatewayHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	log.Println("服务开启，监听端口：", config.GatewayPort)
	return http.ListenAndServe(":"+strconv.Itoa(int(config.GatewayPort)), mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
