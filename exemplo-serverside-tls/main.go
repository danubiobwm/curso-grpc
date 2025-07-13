package main

import (
	"context"
	"log"
	"net"

	"github.com/danubiobwm/curso-grpc/exemplo-serverside-tls/src/pb/products"
	"google.golang.org/grpc"
)

type server struct {
	products.UnimplementedProductServiceServer
}

func (s *server) FindAll(ctx context.Context, req *products.ListProductsRequest) (*products.ListProductResponse, error) {
	productList := make([]*products.Product, 0)
	productList = append(productList, &products.Product{
		Id:    1,
		Title: "Product 1",
	})
	return &products.ListProductResponse{
		Products: productList,
	}, nil
}

func main() {
	srv := server{}
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	products.RegisterProductServiceServer(s, &srv)
	log.Println("Server is running on port 9090")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
