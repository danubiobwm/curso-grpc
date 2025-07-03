package main

import (
	"context"
	"log"

	"github.com/danubiobwm/curso-grpc/api-calc-client/src/pb/calc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := calc.NewCalcServiceClient(conn)
	stream, err := client.Calc(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}
	nums := []int32{1, 2, 3, 4, 5}
	for _, num := range nums {
		if err := stream.Send(&calc.Input{Value: num}); err != nil {
			log.Fatalf("Failed to send number %d: %v", num, err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}
	log.Printf("Sum: %+v", response)

}
