package main

import (
	"context"
	"fmt"
	"grpc-client/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	grpcUser, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcUser.Close()

	client := pb.NewAttendanceServiceClient(grpcUser)
	// request := &pb.CheckInRequest{Username: "Bejo", Datetime: "2024-03-12 07:00"}
	now := time.Now().Format("2006-01-02 15:04")
	request := &pb.CheckInRequest{Username: "Bejo", Datetime: now}

	resp, err := client.CheckIn(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive response => %s ", resp)
}
