package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	pb "github.com/yiting-tom/LF_Golang/grpc/unary/calculator"
	utils "github.com/yiting-tom/LF_Golang/grpc/unary/utils/argParser"
	"google.golang.org/grpc"
)

func main() {
	log.Info("starting gRPC client...")

	// Dial the server
	conn, err := grpc.Dial(
		"localhost:5000",
		grpc.WithInsecure(), // just keep it simple, without auth.
	)
	if err != nil {
		log.Fatal("failed to connect: %v", err)
	}
	defer conn.Close()

	// create a new calculator client
	c := pb.NewCalculatorServiceClient(conn)

	// create a new request
	req, err := utils.ClientArgParse()
	if err != nil {
		log.Fatal("failed to parse args: %v", err)
		log.Info("use default args: 1 + 2")
		req = &pb.FormulaRequest{
			A:        1,
			B:        2,
			Operator: pb.Operator_ADDITION,
		}
	}

	// call the server
	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatal("failed to call server: %v", err)
	}

	// print the response
	log.Info("response: ", res)
}
