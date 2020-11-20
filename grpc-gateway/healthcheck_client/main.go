package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-gateway/healthcheck_client/healthcheck"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := healthcheck.NewHealthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.Check(ctx, &healthcheck.HealthCheckRequest{})
	if err != nil {
		log.Fatalf("Could not check service: %v", err)
	}
	log.Printf("resp: %s", resp.String())
}
