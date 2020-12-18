package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"os"

	"haproxy.com/grpc/rolldice"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	lis, err := net.Listen("tcp", serverAddress)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	rolldice.RegisterRollDiceServiceServer(grpcServer, &rollDiceServer{})
	log.Println("Listening on address ", serverAddress)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type rollDiceServer struct{}

func (s *rollDiceServer) RollDice(ctx context.Context, request *rolldice.RollDiceRequest) (*rolldice.RollDiceResponse, error) {
	die_numbers := []int32{1, 2, 3, 4, 5, 6}
	rolls := []int32{}
	var i int32

	for i = 0; i < request.NumberOfDice; i++ {
		// roll die
		roll := die_numbers[rand.Intn(5)]
		rolls = append(rolls, roll)
	}

	return &rolldice.RollDiceResponse{Rolls: rolls}, nil
}
