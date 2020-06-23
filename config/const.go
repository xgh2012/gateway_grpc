package config

type (
	PortType uint
)

var (
	RpcStatus = false
)

const (
	GrpcPort    PortType = 1090
	GatewayPort PortType = 2090
)
