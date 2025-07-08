package main

import (
	"fmt"
	"io"
	"net"

	"github.com/danubiobwm/curso-grpc/api-shoppingcart/src/pb/shoppingcart"
	"google.golang.org/grpc"
)

type server struct {
	shoppingcart.ShoppingCartServiceServer
}

func (s *server) AddItem(srv shoppingcart.ShoppingCartService_AddItemServer) error {
	var quantityItems int32 = 0
	var priceTotal float64 = 0.0
	for {
		newItem, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				// Quando o cliente termina de enviar, retorna o total
				return srv.Send(&shoppingcart.ShoppingCartTotal{
					QuantityItems: quantityItems,
					TotalPrice:    priceTotal,
				})
			}
			return fmt.Errorf("error receiving item: %v", err)
		}

		// Processa o item recebido
		quantityItems += newItem.GetQuantity()
		priceTotal += newItem.GetPriceUnit() * float64(newItem.GetQuantity())
	}
}

func main() {
	fmt.Println("Shopping Cart Service is running...")
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("Failed to listen on port 9090: %v\n", err)
		return
	}
	s := grpc.NewServer()
	shoppingcart.RegisterShoppingCartServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
		return
	}
	fmt.Println("Server is listening on port 9090")
	defer listener.Close()

}
