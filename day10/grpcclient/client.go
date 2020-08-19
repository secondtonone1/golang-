package main

//client.go
import (
	pb "golang-/day10/customer"
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

//这里是封装的两个函数，也可以不封装直接调用
// createCustomer calls the RPC method CreateCustomer of CustomerServer
func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not Create Consumer%v", err)
	}

	if resp.Success {
		log.Printf("A new Customer has been added with id %d", resp.Id)
	}
}

// getCustomers calls the RPC method GetCustomers of CustomerServer
func getCustomer(client pb.CustomerClient, filter *pb.CustomerFilter) {
	stream, err := client.GetCustomer(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers:%v", err)
	}

	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers() = _ %v", client, err)
			break
		}
		log.Printf("Customer:%v", customer)
	}
}

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect :%v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewCustomerClient(conn)
	customer := &pb.CustomerRequest{
		Id:    101,
		Name:  "Li Lei",
		Email: "Li@163.com",
		Addresses: []*pb.Address{
			&pb.Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: false,
			},
			&pb.Address{
				Street:            "Greenfield",
				City:              "Kochi",
				State:             "KL",
				Zip:               "68356",
				IsShippingAddress: true,
			},
		},
	}
	createCustomer(client, customer)
	customer = &pb.CustomerRequest{
		Id:    102,
		Name:  "Han Meimei",
		Email: "HanMei@163.com",
		Addresses: []*pb.Address{
			&pb.Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	createCustomer(client, customer)
	// Filter with an empty Keyword
	filter := &pb.CustomerFilter{Keyword: "Li Lei"}
	getCustomer(client, filter)
}
