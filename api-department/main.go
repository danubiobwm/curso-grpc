package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/danubiobwm/curso-grpc/api-department/src/pb/department"
	"google.golang.org/grpc"
)

type server struct {
	department.DepartmentServiceServer
}

func (s *server) ListPerson(req *department.ListPersonRequest, srv department.DepartmentService_ListPersonServer) error {
	file, err := os.Open("./data.csv")
	if err != nil {
		return fmt.Errorf("failed to open data file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ";")
		id, _ := strconv.Atoi(data[0])
		name := data[1]
		email := data[2]
		income, _ := strconv.Atoi(data[3])
		departmentId, _ := strconv.Atoi(data[4])

		if int32(departmentId) == req.GetDepartmentId() {
			if err := srv.Send(&department.ListPersonResponse{
				Id:           int32(id),
				Name:         name,
				Email:        email,
				Income:       int32(income),
				DepartmentId: int32(departmentId),
			}); err != nil {
				return fmt.Errorf("failed to send person: %v", err)
			}
		}
	}
	return nil
}

func main() {
	fmt.Println("Starting grpc Department Service...")
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()
	department.RegisterDepartmentServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
		return
	}
	fmt.Println("gRPC Department Service is running on port 9090")
	defer listener.Close()

}
