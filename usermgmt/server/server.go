package main

import (
	"fmt"
	"context"
	"math/rand"
	"net"
	"log"

	pb "github.com/dhruvit2/usermgmt/usermgmt"
	"google.golang.org/grpc"
)

type UserManagementServer struct {
        pb.UnimplementedUserManagementServer
}

func (u *UserManagementServer) CreateNewUser(ctx context.Context, nu *pb.NewUser) (*pb.User, error) {
	fmt.Printf("name %v age %v \n", nu.GetName(), nu.GetAge())

	var user_id  int32 = int32(rand.Intn(1000))
	return &pb.User{Name:nu.GetName(), Age:nu.GetAge(), Id:user_id}, nil
}

func main() {

	lis, err := net.Listen("tcp",":5050")
	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.ServerOption

	s := grpc.NewServer(opts...)

	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("Server listening at %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve the listener\n")
	}
}