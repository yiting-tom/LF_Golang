package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	pb "github.com/yiting-tom/LF_Golang/grpc/route"
	"google.golang.org/grpc"
)

func main() {
	// Dial is a convenience function for dialing
	// a client connection to a server.
	// the first argument is the target server address,
	// and others are dial options.
	conn, err := grpc.Dial(
		":5000",
		grpc.WithInsecure(), // Since we don't have server yet
		grpc.WithBlock(),    // Block until the connection is established
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	feature, err := client.GetFeature(
		context.Background(),
		&pb.Point{
			Latitude:  37,
			Longitude: 122,
		},
	)
	if err != nil {
		log.Fatalf("failed to get feature: %v", err)
	}

	fmt.Println(feature)
}
