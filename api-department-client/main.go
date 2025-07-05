package main

import (
	"context"
	"io"
	"log"

	"github.com/danubiobwm/curso-grpc/api-department-client/src/pb/department"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := department.NewDepartmentServiceClient(conn)
	stream, err := client.ListPerson(context.Background(), &department.ListPersonRequest{DepartmentId: 2})
	if err != nil {
		log.Fatalf("Error calling ListPerson: %v", err)
	}

	for {
		person, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving person: %v", err)
		}
		log.Printf("Received person: %s", person)
	}

}
