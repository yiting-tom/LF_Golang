package main

import (
	"context"
	"net"

	log "github.com/sirupsen/logrus"
	pb "github.com/yiting-tom/LF_Golang/grpc/unary/calculator"
	"google.golang.org/grpc"
)

// calculatorServer wraps the CalculatorServiceServer interface
type calculatorServer struct {
	pb.UnimplementedCalculatorServiceServer
}

// Calculate implements the Calculate method of the CalculatorServiceServer interface
func (s *calculatorServer) Calculate(ctx context.Context, req *pb.FormulaRequest) (*pb.SolutionResponse, error) {
	log.Info("Calculate() called with request: ", req)

	var sol int32
	a := req.GetA()
	b := req.GetB()

	switch req.GetOperator() {
	case pb.Operator_ADDITION:
		sol = a + b
	case pb.Operator_SUBTRACTION:
		sol = a - b
	case pb.Operator_MULTIPLICATION:
		sol = a * b
	case pb.Operator_DIVISION:
		sol = a / b
	default:
		res := &pb.SolutionResponse{
			Response: &pb.SolutionResponse_Error{
				Error: "unsupported operator",
			},
		}
		return res, nil
	}

	res := &pb.SolutionResponse{
		Response: &pb.SolutionResponse_Solution{
			Solution: sol,
		},
	}

	return res, nil
}

func main() {
	log.Info("starting gRPC server...")

	// listen on port 5000
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the calculator service with the gRPC server
	pb.RegisterCalculatorServiceServer(grpcServer, &calculatorServer{})

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
