package main

import (
	"context"
	"gateway_grpc/config"
	pb "gateway_grpc/protoc"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	//连接服务端句柄
	//WithInsecure()不指定安全选项
	conn, err := grpc.Dial("localhost:"+strconv.Itoa(int(config.GrpcPort)), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewGatewayEchoClient(conn)

	name := "World XGH"
	//获取命令行参数
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	for i := 1; i <= 10; i++ {
		name = "World XGH " + strconv.Itoa(i)

		go func(name string) {
			//设置上下文超时
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			//响应
			resp, err := cli.Echo(ctx, &pb.EchoRequest{Name: name})
			if err != nil {
				log.Fatalf("could not succ: %v", err)
			}
			log.Printf("Receive from server: %s", resp.Message)
		}(name)

		time.Sleep(500 * time.Millisecond)
	}
}
