package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"sync"
	"time"

	"google.golang.org/grpc"
	_ "net/http/pprof"
	ps "netangels/piservice/proto"
)

var client1 ps.CalcPiClient
var client2 ps.CalcPiClient
var client3 ps.CalcPiClient

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New request")
	start := time.Now()
	var wg sync.WaitGroup
	var err error

	//sample := os.Args[1]
	//n, err := strconv.Atoi(sample)
	n := 1000

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

	s := n / 3
	a := 2 * n / 3
	go func(req *ps.PiRequest) {
		defer wg.Done()
		if resp2, err = client2.GeneratePi(context.Background(), req); err != nil {
			log.Fatalf("Could not get answer: %v", err)
		}

	}(&ps.PiRequest{Start: int32(s), Accuracy: int32(a)})

	wg.Add(1)

	s = 2 * n / 3
	a = n
	go func(req *ps.PiRequest) {
		defer wg.Done()
		if resp3, err = client3.GeneratePi(context.Background(), req); err != nil {
			log.Fatalf("Could not get answer: %v", err)
		}

	}(&ps.PiRequest{Start: int32(s), Accuracy: int32(a)})

	wg.Wait()

	log.Println("Pi:", resp1.Pi+resp2.Pi+resp3.Pi)

	if _, err = w.Write([]byte(fmt.Sprintf("%v", resp1.Pi+resp2.Pi+resp3.Pi))); err != nil {
		log.Println(err)
	}

	duration := time.Since(start)
	log.Printf("Duration: %v", duration)
}

func main() {

	conn1, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	conn2, _ := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	conn3, _ := grpc.Dial("127.0.0.1:8083", grpc.WithInsecure())

	client1 = ps.NewCalcPiClient(conn1)
	client2 = ps.NewCalcPiClient(conn2)
	client3 = ps.NewCalcPiClient(conn3)

	r := http.NewServeMux()

	// Установить маршрут доступа
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	r.HandleFunc("/hello", myHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:9090", nil))

}
