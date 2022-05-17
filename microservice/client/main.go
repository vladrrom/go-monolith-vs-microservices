package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	ps "netangels/piservice/proto"

	"google.golang.org/grpc"
)

func main() {

	var wg sync.WaitGroup

	conn1, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	conn2, _ := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	conn3, _ := grpc.Dial("127.0.0.1:8083", grpc.WithInsecure())

	start := time.Now()
	client1 := ps.NewCalcPiClient(conn1)
	client2 := ps.NewCalcPiClient(conn2)
	client3 := ps.NewCalcPiClient(conn3)

	var err error

	sample := os.Args[1]
	n, err := strconv.Atoi(sample)

	log.Printf("Entered quantity of goroutines: %v", n)

	var resp1 *ps.PiResponse
	var resp2 *ps.PiResponse
	var resp3 *ps.PiResponse

	wg.Add(1)

	go func(req *ps.PiRequest) {
		defer wg.Done()
		if resp1, err = client1.GeneratePi(context.Background(), req); err != nil {
			log.Fatalf("Could not get answer: %v", err)
		}
	}(&ps.PiRequest{Start: int32(0), Accuracy: int32(n / 3)})

	wg.Add(1)

	go func(req *ps.PiRequest) {
		defer wg.Done()
		if resp2, err = client2.GeneratePi(context.Background(), req); err != nil {
			log.Fatalf("Could not get answer: %v", err)
		}

	}(&ps.PiRequest{Start: int32(n / 3), Accuracy: int32((2 / 3) * n)})

	wg.Add(1)

	go func(req *ps.PiRequest) {
		defer wg.Done()
		if resp3, err = client3.GeneratePi(context.Background(), req); err != nil {
			log.Fatalf("Could not get answer: %v", err)
		}

	}(&ps.PiRequest{Start: int32((2 / 3) * n), Accuracy: int32(n)})

	wg.Wait()

	log.Println("Pi:", resp1.Pi+resp2.Pi+resp3.Pi)

	duration := time.Since(start)
	log.Printf("Duration: %v", duration)
}
