package main

import (
	"context"
	"fmt"
	"log"
	"pie_fire_dire_grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewBeefServiceClient(conn)

	r, err := c.GetBeefSummary(context.Background(), &proto.Empty{})
	if err != nil {
		log.Fatalf("could not get beef summary: %v", err)
	}

	fmt.Println("Meat Count:", r.Meats)
}
