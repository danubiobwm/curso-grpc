package main

import (
	"fmt"
	"io"
	"net"

	"github.com/danubiobwm/curso-grpc/api-calc/src/pb/calc"
	"google.golang.org/grpc"
)

type server struct {
	calc.CalcServiceClient
	calc.UnimplementedCalcServiceServer
}

func (s *server) Calc(stream calc.CalcService_CalcServer) error {
	var quantity int32 = 0
	var total int32 = 0

	for {
		input, err := stream.Recv()
		if err == io.EOF {
			avg := float64(total) / float64(quantity)
			return stream.SendAndClose(&calc.Output{
				Quantity: quantity,
				Aver:     avg,
				Total:    total,
			})

		}
		if err != nil {
			return fmt.Errorf("error receiving data: %v", err)
		}
		quantity++
		total += input.GetValue()
		fmt.Printf("Received value: %d, Total: %d, Quantity: %d\n", input.GetValue(), total, quantity)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	calc.RegisterCalcServiceServer(grpcServer, &server{})

	fmt.Println("Server is running on port 9090")
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
		return
	}

}
