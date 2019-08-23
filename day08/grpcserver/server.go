// consignment-service/main.go
package main

import (
	"log"
	"net"

	// 导入生成的 protobuf 代码
	pb "golang-/day08/consignment"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - 一个模拟数据存储的虚拟仓库，以后我们会替换成真实的数据仓库
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// 服务需要实现所有在 protobuf 里定义的方法。
// 你可以参考 protobuf 生成的 go 文件中的接口信息。
type service struct {
	repo IRepository
}

// CreateConsignment - 目前只创建了这个方法，包括 `ctx` (环境信息)和 `req`(委托请求)两个参数，会通过 gRPC 服务器进行处理
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	// 保存委托
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// 返回和 protobuf 中定义匹配的 `Response` 消息
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {

	repo := &Repository{}

	// 启动 gRPC 服务器。
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	s := grpc.NewServer()

	// 注册服务到 gRPC 服务器，会把已定义的 protobuf 与自动生成的代码接口进行绑定。
	pb.RegisterShippingServiceServer(s, &service{repo})

	// 在 gRPC 服务器上注册 reflection 服务。
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
