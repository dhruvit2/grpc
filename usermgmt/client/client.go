package main

import (
	"log"
	"context"
	"time"
	"fmt"
	"io"

	pb "github.com/dhruvit2/usermgmt/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5050"
)


func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Dial failed %v \n", err)
	}

	defer conn.Close()
	client := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	var newUser = make(map[string]int32)
	newUser["dhr"] = 30
	newUser["dhr1"] = 31

	for name, age := range newUser {
		user, err := client.CreateNewUser(ctx, &pb.NewUser{Name: name, Age:age})
		if err != nil {
			log.Fatal("COuld not create user %v \n", err)
		}

		log.Printf("User details Name %v Age %v Id %v \n", user.GetName(), user.GetAge(), user.GetId())
	}

	resStream, err := client.GreetUser(ctx, &pb.NewUser{Name: "ser strea", Age:67})
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		fmt.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}
}
