package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"runtime"
	"fmt"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := NewStreamingServiceClient(conn)

	r, err := c.StreamIt(context.Background(), &Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for {
		payload, err := r.Recv()
		if err != nil {
			fmt.Println()
			log.Printf("Error received %+v\n", err)
			break
		}
		_ = payload
		ms := runtime.MemStats{}
		runtime.ReadMemStats(&ms)
		fmt.Printf("\rServing. Alloc %10d, Sys %10d.", ms.Alloc, ms.Sys)
	}
}
