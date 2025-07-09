package main

import (
	"context"
	"fmt"
	"log"

	"github.com/danubiobwm/curso-grpc/api-shoppingcart-client/src/pb/shoppingcart"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := shoppingcart.NewShoppingCartServiceClient(conn)
	stream, err := client.AddItem(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	waitch := make(chan struct{})

	go func() {
		for {
			response, err := stream.Recv()
			if err != nil {
				close(waitch)
				return
			}
			if err != nil {
				log.Fatalf("Error receiving response:", err)
			}
			fmt.Printf("Item added to cart:%+v\n", response)
		}
	}()

	items := []shoppingcart.AddProduct{
		{ProductId: 1, Quantity: 2, PriceUnit: 2.6},
		{ProductId: 2, Quantity: 4, PriceUnit: 1.7},
		{ProductId: 3, Quantity: 6, PriceUnit: 3.8},
	}

	for _, v := range items {
		if err := stream.Send(&v); err != nil {
			log.Fatalf("Error sending item: %v", err)
		}
		fmt.Printf("-> sent: %+v\n", v)
	}

	stream.CloseSend()

	<-waitch

}
