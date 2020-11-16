package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "productinfo/client/ecommerce"
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
	c := pb.NewProductInfoClient(conn)

	//for test
	name := "Apple iPhone 12"
	description := `Test`
	//price := float32(1000.0)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//insert
	r, err := c.AddProduct(ctx, &pb.Product{
		Name:        name,
		Description: description,
	})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	//read
	product, err := c.GetProduct(ctx, &pb.ProductID{Value:r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: ", product.String())
}
