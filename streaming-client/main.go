package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	pb "go-micro-1/pb"

	"google.golang.org/grpc"
)

func main() {
	var (
		server   = flag.String("server", "127.0.0.1:10000", "Server address.")
		name     = flag.String("name", "", "Username to use")
		lastname = flag.String("lastname", "", "Lastname to use")
	)

	flag.Parse()

	conn, err := grpc.Dial(*server, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	stream, err := c.SayHelloStream(context.Background(), &pb.Request{Name: *name, Lastname: *lastname})

	if err != nil {
		log.Fatal(err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response.Message)
	}
}
