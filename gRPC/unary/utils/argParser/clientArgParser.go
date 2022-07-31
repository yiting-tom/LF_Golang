package utils

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	pb "github.com/yiting-tom/LF_Golang/grpc/unary/calculator"
)

var (
	operands [2]int
	operator string
	Atoo     = map[string]pb.Operator{
		"+": pb.Operator_ADDITION,
		"-": pb.Operator_SUBTRACTION,
		"*": pb.Operator_MULTIPLICATION,
		"/": pb.Operator_DIVISION,
	}
)

func init() {
	flag.IntVar(&operands[0], "a", 1, "operand a")
	flag.IntVar(&operands[1], "b", 2, "operand b")
	flag.StringVar(&operator, "op", "+", "operator")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
	flag.PrintDefaults()
}

func ClientArgParse() (*pb.FormulaRequest, error) {
	flag.Parse()

	args := flag.Args()

	if len(args) != 3 {
		return nil, fmt.Errorf("invalid number of arguments")
	}

	a, _ := strconv.Atoi(args[0])
	op := Atoo[args[1]]
	b, _ := strconv.Atoi(args[2])

	return &pb.FormulaRequest{
		A:        int32(a),
		B:        int32(b),
		Operator: op,
	}, nil
}
