# gateway_grpc

1、安装包 
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

2、生成service
  window 生成service
  protoc --proto_path=../ -I/m/goProgram/bin -I. -I/m/goProgram/src -I/m/goProgram/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. *.proto
  
  mac 生成service
  protoc -I/Users/xhome/go/bin -I. -I/Users/xhome/go/src -I/Users/xhome/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. *.proto
  
3、gw 生成
   window gw 生成  
   protoc --proto_path=../ -I/m/goProgram/bin -I. -I/m/goProgram/src -I/m/goProgram/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. *.proto
   
   mac 生成 gw
   protoc -I. -I/Users/xhome/go/src -I/Users/xhome/go/bin -I/Users/xhome/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. gateway/echo.proto
   
   
   表示googleapis 包所在路径 ： -I/m/goProgram/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis


4、其他语言调用参数
    {"header":{"client":"xxx"},"params":[{"key":"xxx1","value":"yyy1"},{"key":"xxx2","value":"yyy2"},{"key":"xxx3","value":"yyy3"}],"mobile":"test"}