package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb_address "go-protobuf/pb/addressbook"
	pb "go-protobuf/pb/userInfo"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type pbUserInfoserver struct {
	pb.UnimplementedUserInfoServer
}

type pbAddressServer struct {
	pb_address.UnimplementedAddressInfoServer
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserInfoServer(s, &pbUserInfoserver{})
	pb_address.RegisterAddressInfoServer(s, &pbAddressServer{})
	log.Printf("server listening at: %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *pbUserInfoserver) GetUserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoRes, error) {
	fmt.Println("GetUserInfo server recv:", req.UserId)
	return &pb.UserInfoRes{
		Name:    "Superman",
		Gender:  pb.UserInfoRes_FEMALE,
		Habbits: []string{"Baseball", "Movie", "Hiking"},
	}, nil
}

func (s *pbAddressServer) GetAddressInfo(ctx context.Context, req *pb_address.Person) (*pb_address.Person, error) {
	fmt.Println("GetAddressInfo server recv:", req.Name)
	return &pb_address.Person{
		Name:  req.Name,
		Email: "qwe@gmail.com",
		Phones: []*pb_address.Person_PhoneNumber{
			{Number: "123", Type: pb_address.Person_MOBILE},
			{Number: "456", Type: pb_address.Person_HOME},
			{Number: "789", Type: pb_address.Person_WORK},
		},
	}, nil
}
