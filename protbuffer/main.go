package main

import (
	"fmt"
	"os"

	"github.com/danubiobwm/curso-grpc/protbuffer/src/pb/users"
	"google.golang.org/protobuf/proto"
)

func main() {
	//createData()
	readData()
}

func readData() {
	var user users.User
	data, err := os.ReadFile("user.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	err = proto.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return
	}
	fmt.Printf("User: %+v\n", user)
}

func createData() {
	user := users.User{
		Id:       1,
		Name:     "Danubio",
		Email:    "jotest@bol.com.br",
		Password: "123456",
	}
	out, err := proto.Marshal(&user)
	if err != nil {
		fmt.Println("Error marshalling user:", err)
	}
	err = os.WriteFile("user.txt", out, 0644)

	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
