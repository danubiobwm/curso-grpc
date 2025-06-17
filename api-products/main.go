package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/danubiobwm/curso-grpc/api-products/src/pb/products"
	"google.golang.org/grpc"
)

type server struct {
	products.ProductServiceServer
}

func (s *server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {
	return &products.Product{}, nil
}

func (s *server) FindAll(ctx context.Context, product *products.Product) (*products.ProductList, error) {
	fmt.Println("FindAll called with product:", product)
	return &products.ProductList{}, nil
}

func main() {
	fmt.Println("Starting gRPC server...")
	srv := server{}

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen:", err)
	}

	s := grpc.NewServer()

	products.RegisterProductServiceServer(s, &srv)
	log.Println("Server is running on port 9090")

	if err := s.Serve(listener); err != nil {
		log.Fatalf("error on server. error:", err)
	}

}
