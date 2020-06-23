package main

import (
	pb "go-micro-1/pb"
	"log"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{Message: "Hello " + request.Name + request.Lastname}, nil
}

func (s *server) SayHelloStream(request *pb.Request, stream pb.Greeter_SayHelloStreamServer) error {
	for {
		err := stream.Send(&pb.Response{Message: "Hello " + request.Name + " " + request.Lastname})
		if err != nil {
			log.Println(err)
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	log.Println("hello world service starting...")
	ln, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("Helloworld service start successfully")
	s.Serve(ln)
}
