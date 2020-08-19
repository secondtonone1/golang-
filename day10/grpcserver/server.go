package main

// server.go
import (
	"context"
	pb "golang-/day10/customer"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	savedCustomers []*pb.CustomerRequest
}

func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)
	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

func (s *server) GetCustomer(filter *pb.CustomerFilter, stream pb.Customer_GetCustomerServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}

			if err := stream.Send(customer); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterCustomerServer(s, &server{})
	s.Serve(lis)
}
