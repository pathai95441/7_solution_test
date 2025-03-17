package main

import (
	"context"
	"log"
	"net"

	"pie_fire_dire_grpc/handler"
	"pie_fire_dire_grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedBeefServiceServer
}

func (s *server) GetBeefSummary(ctx context.Context, in *proto.Empty) (*proto.BeefSummary, error) {
	meatCount, err := handler.BeefSummaryHandler()
	if err != nil {
		return nil, err
	}

	beefSummary := &proto.BeefSummary{Meats: *meatCount}
	return beefSummary, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterBeefServiceServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
