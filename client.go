package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"haproxy.com/grpc/rolldice"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := os.Getenv("SERVER_ADDRESS")
	numberOfDice32, err := strconv.ParseInt(os.Getenv("NUMBER_OF_DICE"), 10, 32)

	if err != nil {
		log.Fatalf("Could not parse NUMBER_OF_DICE: %v", err)
	}

	// convert int32 to int
	numberOfDice := int32(numberOfDice32)

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	defer conn.Close()
	client := rolldice.NewRollDiceServiceClient(conn)
	ctx := context.Background()

	for {
		result, err := client.RollDice(ctx, &rolldice.RollDiceRequest{NumberOfDice: numberOfDice})

		if err != nil {
			log.Fatalf("Could not get result from RollDice: %v", err)
		}

		log.Printf("Rolling %d dice...", numberOfDice)

		for _, roll := range result.Rolls {
			log.Printf("roll: %d", roll)
		}

		time.Sleep(2 * time.Second)
	}
}
