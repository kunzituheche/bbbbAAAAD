package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	proto "MESSAGE-SERVICE/gen/proto/go/mes/v1"
)

//模拟短信模块服务端代码

type MessageServer struct {
	proto.UnimplementedMessageServer
}

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8081", "gRPC server endpoint")
)

const (
	HTTPLISTEN = ":8081"
	GRPCLISTEN = ":9090"
)

func (m *MessageServer) Message(ctx context.Context, in *proto.MessageRequest) (*proto.MessageResponse, error) {
	//通过传来的request中的phone和验证码查询数据库中是否一致
	//这里写MessageService中的方法
	return &proto.MessageResponse{
		Message: "收到号码" + in.Phone,
		Code:    "2",
	}, nil
}

func (m *MessageServer) Message_http(ctx context.Context, in *proto.MessageRequest) (*proto.MessageResponse, error) {
	return &proto.MessageResponse{
		Message: "grpc-gateway测试！！",
		Code:    "3",
	}, nil
}

//同时启动grpc和http服务
func HttpAndGrpc() {
	lis, err := net.Listen("tcp", GRPCLISTEN) //监听grpc 9090端口
	if err != nil {
		panic(err.Error())
	}
	grpcServer := grpc.NewServer()
	proto.RegisterMessageServer(grpcServer, &MessageServer{})
	go grpcServer.Serve(lis) //启动grpc服务

	mux := runtime.NewServeMux()
	err = proto.RegisterMessageHandlerFromEndpoint(context.Background(), mux, GRPCLISTEN, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		panic(err.Error())
	}
	http.ListenAndServe(HTTPLISTEN, mux) //http 8081端口
}

func HttpServerRun() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	proto.RegisterMessageHandlerFromEndpoint(ctx, nil, *grpcServerEndpoint, opts)

	return http.ListenAndServe(":8081", mux)
}

func ServerRun() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer() // 创建gRPC服务器

	// 为 message 服务注册业务实现（将 message 服务绑定到 GRPC 服务容器上）
	proto.RegisterMessageServer(grpcServer, &MessageServer{})

	// 注册反射服务，这个服务是 CLI 使用的，跟服务本身没有关系
	//reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}

}
