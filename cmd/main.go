package main

import (
	cube_client "example/microabstraction/external/client/cube"
	square_client "example/microabstraction/external/client/square"
	"example/microabstraction/internal/server"
	manager "example/microabstraction/internal/server/manager"
	"example/microabstraction/internal/service"
	"log"
)

func main() {
	cubeClient, err := cube_client.NewCubeRestfulClient()
	if err != nil {
		log.Fatal(err)
	}
	squareClient, err := square_client.NewSquareGrpcClient()
	if err != nil {
		log.Fatal(err)
	}
	sumService := service.NewSumService(squareClient, cubeClient)
	srv := server.NewRestfulServer(sumService)
	manager, err := manager.NewServerManager(srv)
	if err != nil {
		log.Fatal(err)
	}
	manager.StartServer()
}
