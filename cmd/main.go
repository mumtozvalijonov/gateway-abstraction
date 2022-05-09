package main

import (
	"example/microabstraction/api/server"
	"example/microabstraction/internal/service"
)

func main() {
	sumService := service.NewSumService()
	server := server.NewGrpcServer(sumService)
	server.Run()
}
