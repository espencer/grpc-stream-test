package main

import (
	"fmt"
	"log"
	"net"
	"math/rand"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"

	events = 1000 * 1000 * 1000

	payload_size = 1024
)

// server is used to implement StreamingServiceClient
type server struct{}

func (s *server) StreamIt(in *Empty, downstream StreamingService_StreamItServer) error {

	payload := RandStringBytes(payload_size)

	for i := 0; i < events; i++ {
		ms := runtime.MemStats{}
		runtime.ReadMemStats(&ms)
		fmt.Printf("\rServing. Alloc %10d, Sys %10d.", ms.Alloc, ms.Sys)

		err := downstream.Send(&Event{payload})
		if err != nil {
			fmt.Println()
			log.Fatalf("error sending event: %v", err)
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterStreamingServiceServer(s, &server{})
	reflection.Register(s)
	fmt.Printf("Ready.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
