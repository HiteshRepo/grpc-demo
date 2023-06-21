package grpcapp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/hiteshrepo/grpc-demo/internal/pkg/model"
	"github.com/hiteshrepo/grpc-demo/internal/pkg/proto"
	"github.com/hiteshrepo/grpc-demo/internal/pkg/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type GrpcApp struct {
	proto.UnimplementedBookServiceServer
	bookRepo repo.BookRepo
}

func NewGrpcApp() GrpcApp {
	return GrpcApp{bookRepo: repo.NewBookRepo()}
}

func (a *GrpcApp) Start() {
	fmt.Println("Starting book server")

	port := "50051"
	if len(os.Getenv("GRPC_SERVER_PORT")) > 0 {
		port = os.Getenv("GRPC_SERVER_PORT")
	}

	servAddr := fmt.Sprintf("0.0.0.0:%s", port)

	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: time.Minute * 5,
		}),
	}
	s := grpc.NewServer(opts...)
	proto.RegisterBookServiceServer(s, a)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *GrpcApp) Shutdown() {

}

func (a *GrpcApp) AddBook(_ context.Context, req *proto.BookRequest) (*proto.Empty, error) {
	book := &model.Book{
		ISBN:   int(req.Isbn),
		Author: req.Author,
		Name:   req.Name,
	}

	a.bookRepo.StoreBook(book)

	return &proto.Empty{}, nil
}

func (a *GrpcApp) ListBooks(ctx context.Context, _ *proto.Empty) (*proto.BooksResponse, error) {

	books := a.bookRepo.GetBooks()
	b, err := json.Marshal(books)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling books", err.Error())
	}

	pbBooks := []*proto.BookRequest{}
	err = json.Unmarshal(b, &pbBooks)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling books", err.Error())
	}

	return &proto.BooksResponse{Books: pbBooks}, nil
}
