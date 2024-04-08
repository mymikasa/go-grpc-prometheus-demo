package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"log"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	pb "github.com/mymikasa/echo/protobuf/gen/go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

var (
	conn *grpc.ClientConn
)

func main() {
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", os.Getenv("GRPC_SERVER_HOST"), os.Getenv("GRPC_PORT")),
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(time.Second),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
	)

	if err != nil {
		log.Printf("Connection error: %v", err)
	}

	defer conn.Close()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/echo", echoHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("GRPC_CLIENT_PORT")), nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	client := pb.NewEchoClient(conn)

	message := &pb.EchoRequest{
		Message: r.URL.Query().Get("m"),
	}

	response, err := client.Echo(context.Background(), message)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	fmt.Fprintf(w, "Response: %v", response.Message)

}
