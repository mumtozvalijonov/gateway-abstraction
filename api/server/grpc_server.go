package server

import (
	"context"
	"example/microabstraction/internal/service"
	"flag"
	"fmt"
	"log"
	"net"

	pb "example/microabstraction/api/server/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type GrpcServer struct {
	sumService *service.SumService
}

type server struct {
	sumService *service.SumService
	pb.UnimplementedSumServiceServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	sum := s.sumService.GetSum(in.A, in.B)
	return &pb.SumResponse{Result: sum}, nil
}

func (s GrpcServer) Run() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterSumServiceServer(grpcSrv, &server{sumService: s.sumService})

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func NewGrpcServer(service *service.SumService) *GrpcServer {
	return &GrpcServer{sumService: service}
}
