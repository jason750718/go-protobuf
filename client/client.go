package main

import (
	"context"
	"log"
	"time"

	pb_address "go-protobuf/pb/addressbook"
	pb "go-protobuf/pb/userInfo"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("connect to grpc server fail: %v", err)
	}
	defer conn.Close()

	c_userInfo := pb.NewUserInfoClient(conn)
	c_address := pb_address.NewAddressInfoClient(conn)
	sendUserInfo(c_userInfo, "fakeuserid_12313123")
	sendAddress(c_address, "Superwoman")
}

func sendUserInfo(c pb.UserInfoClient, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.GetUserInfo(ctx, &pb.UserInfoReq{UserId: msg})
	if err != nil {
		log.Fatalf("failed to send %v: ", err)
	}
	log.Printf("sendUserInfo response res.Name: %s", res.Name)
	log.Printf("sendUserInfo response res.Gender: %d", res.Gender)
	log.Printf("sendUserInfo response res.Habbits: %v", res.Habbits)
}

func sendAddress(c pb_address.AddressInfoClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.GetAddressInfo(ctx, &pb_address.Person{Name: name})
	if err != nil {
		log.Fatalf("failed to send %v: ", err)
	}
	log.Printf("sendAddress response res.Name: %s", res.Name)
	log.Printf("sendAddress response res.Email: %s", res.Email)
	log.Printf("gRPsendAddressC response res.Phones: %v", res.Phones[0].Type)
}
