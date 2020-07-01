package gateway_service

import (
	"flag"
	"gateway_grpc/config"
	"gateway_grpc/gateway_service/client"
	"log"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:"+strconv.Itoa(int(config.GrpcPort)), "endpoint of Gateway")
)

func run() error {
	var err error
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = client.SetClientRouter(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	log.Println("Http 服务开启，监听端口：", config.GatewayPort)
	return http.ListenAndServe(":"+strconv.Itoa(int(config.GatewayPort)), mux)
}

func Run() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
