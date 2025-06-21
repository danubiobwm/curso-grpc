package main

import (
	"context"
	"fmt"
	"log"

	"github.com/danubiobwm/curso-grpc/api-products-client/src/pb/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(
		insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error while connecting to server:", err)
	}
	defer conn.Close()

	FindAllProducts(conn)
	//CreateProduct(conn)

}

func FindAllProducts(conn *grpc.ClientConn) {
	productClient := products.NewProductServiceClient(conn)
	productList, err := productClient.FindAll(context.Background(), &products.Product{})
	if err != nil {
		log.Fatalf("Error while calling FindAll: %v", err)
	}

	fmt.Printf("Products found: %+v\n", productList)
}

func CreateProduct(conn *grpc.ClientConn) {
	productClient := products.NewProductServiceClient(conn)
	newProduct := &products.Product{
		Id:          1,
		Name:        "Celular Nokia",
		Description: "Description of Product 2",
		Price:       10.99,
		Quantity:    100,
	}

	response, err := productClient.Create(context.Background(), newProduct)
	if err != nil {
		log.Fatalf("Error while creating product: %v", err)
	}

	fmt.Printf("Product created with ID: %d\n", response.Id)
}
