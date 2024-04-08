package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	pb "github.com/mymikasa/echo/protobuf/gen/go"
	"google.golang.org/grpc"
)

type grpcServer struct{}

// Echo implements protobuf.EchoServer.
func (g *grpcServer) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: fmt.Sprintf("Hello, World! %s", in.Message)}, nil
}

func main() {
	s, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_SERVER_PORT")))

	if err != nil {
		log.Fatal(err)
	}

	svc := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	)

	pb.RegisterEchoServer(svc, newGrpcSvc())
	grpc_prometheus.Register(svc)
	svc.Serve(s)
}

func newGrpcSvc() *grpcServer {
	return &grpcServer{}
}
