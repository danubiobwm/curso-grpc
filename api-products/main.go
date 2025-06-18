package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/danubiobwm/curso-grpc/api-products/src/pb/products"
	"github.com/danubiobwm/curso-grpc/api-products/src/repository"
	"google.golang.org/grpc"
)

type server struct {
	products.ProductServiceServer
	productRepo *repository.ProductRepository
}

func (s *server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {
	newProduct, err := s.productRepo.Create(*product)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return nil, fmt.Errorf("error creating product: %v", err)
	}

	return &newProduct, nil
}

func (s *server) FindAll(ctx context.Context, product *products.Product) (*products.ProductList, error) {
	productList, err := s.productRepo.FildAll()
	if err != nil {
		log.Printf("Error finding products: %v", err)
		return nil, fmt.Errorf("error finding products: %v", err)
	}
	return &productList, nil
}

func main() {
	fmt.Println("Starting gRPC server...")
	srv := server{productRepo: &repository.ProductRepository{}}

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
