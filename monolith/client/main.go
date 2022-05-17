package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	ps "netangels/piservice/proto"

	"google.golang.org/grpc"
)

var client1 ps.CalcPiClient

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New request")
	start := time.Now()
	var err error

	//sample := os.Args[1]
	//n, err := strconv.Atoi(sample)
	n := 1000

	log.Printf("Entered quantity of goroutines: %v", n)

	var resp1 *ps.PiResponse

	if resp1, err = client1.GeneratePi(context.Background(), &ps.PiRequest{Accuracy: int32(n)}); err != nil {
		log.Fatalf("Could not get answer: %v", err)
	}

	log.Println("Pi:", resp1.Pi)

	if _, err = w.Write([]byte(fmt.Sprintf("%v", resp1.Pi))); err != nil {
		log.Println(err)
	}

	duration := time.Since(start)
	log.Printf("Duration: %v", duration)
}

func main() {

	conn, _ := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
	client1 = ps.NewCalcPiClient(conn)

	//r := http.NewServeMux()
	//
	//// Установить маршрут доступа
	//r.HandleFunc("/debug/pprof/", pprof.Index)
	//r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.HandleFunc("/hello", myHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:9090", nil))
}
