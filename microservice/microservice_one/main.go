package main

import (
	"context"
	"log"
	"math"
	"net"

	ps "netangels/piservice/proto"

	"google.golang.org/grpc"
)

// Функция pi() запускает n горутин для
// вычисления приближения числа pi.
func pi(s, n int32) float64 {
	ch := make(chan float64)

	log.Printf("s:%s n:%s", s, n)

	for k := s; k < n; k++ {
		go term(ch, float64(k))
	}

	f := 0.0

	for k := s; k < n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

type CalcPi struct {
	ps.UnimplementedCalcPiServer
}

func (s *CalcPi) GeneratePi(ctx context.Context, req *ps.PiRequest) (*ps.PiResponse, error) {
	var err error
	response := new(ps.PiResponse)

	response.Pi = pi(req.Start, req.Accuracy)

	return response, err
}

//Type does not implement 'CalcPiServer' as some methods are missing: mustEmbedUnimplementedCalcPiServer()

func main() {
	server := grpc.NewServer()

	instance := new(CalcPi)

	ps.RegisterCalcPiServer(server, instance)

	listener, err := net.Listen("tcp", ":8080")

	//log.Println("AAAAAAAAAAAAAAAA")

	if err != nil {
		log.Fatal("Unable to create gRPC listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
