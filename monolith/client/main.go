package main

import (
	"context"
	"log"
	"time"

	ps "netangels/piservice/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, _ := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
	start := time.Now()
	client := ps.NewCalcPiClient(conn)

	//sample := os.Args[1]
	//n, err := strconv.Atoi(sample)

	n := 50000

	log.Printf("Entered quantity of goroutines: %v", n)

	resp, err := client.GeneratePi(context.Background(),
		&ps.PiRequest{Accuracy: int32(n)})

	if err != nil {
		log.Fatalf("Could not get answer: %v", err)
	}
	log.Println("Pi:", resp.Pi)

	duration := time.Since(start)
	log.Printf("Duration: %v", duration)
}
