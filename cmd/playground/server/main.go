package main

import "mekanicum/server"

type CalculatorServer struct {
	server *server.Server
}

func main() {
	srv := server.New(
		"calculator",
		"0.1.0",
	)

	calculatorSrv := CalculatorServer{
		server: srv,
	}
}
