package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	user "google.golang.org/grpc/grpc-demo/pb"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type UserService struct {
}

func (us *UserService) UserIndex(ctx context.Context, in *user.UserIndexRequest) (*user.UserIndexResponse, error) {
	log.Printf("receive user index request: page %d page_size %d", in.Page, in.PageSize)
	return &user.UserIndexResponse{
		Err: 0,
		Msg: "success",
		Data: []*user.UserEntity{
			{Name: "u1", Age: 11},
			{Name: "u2", Age: 12},
		},
	}, nil
}

func (userService *UserService) UserView(ctx context.Context, in *user.UserViewRequest) (*user.UserViewResponse, error) {
	log.Printf("receive user view request: uid %d", in.Uid)

	return &user.UserViewResponse{
		Err:  0,
		Msg:  "success",
		Data: &user.UserEntity{Name: "james", Age: 28},
	}, nil
}

func (userService *UserService) UserPost(ctx context.Context, in *user.UserPostRequest) (*user.UserPostResponse, error) {
	log.Printf("receive user post request: name %s password %s age %d", in.Name, in.Password, in.Age)

	return &user.UserPostResponse{
		Err: 0,
		Msg: "success",
	}, nil
}

func (userService *UserService) UserDelete(ctx context.Context, in *user.UserDeleteRequest) (*user.UserDeleteResponse, error) {
	log.Printf("receive user delete request: uid %d", in.Uid)

	return &user.UserDeleteResponse{
		Err: 0,
		Msg: "success",
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	user.RegisterUserServer(grpcServer, &UserService{})

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to server: %v", err)
	}
}
