package main

import (
	"log"
	"time"

	pb "github.com/vinimazzarolo/go-grpc-demo/proto"
)

func (s *helloServer) SayHelloServerSideStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerSideStreamingServer) error {
	log.Printf("got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
