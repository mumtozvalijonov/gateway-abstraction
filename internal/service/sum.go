package service

import (
	cube_client "example/microabstraction/external/client/cube"
	square_client "example/microabstraction/external/client/square"
)

type SumService struct {
	squareClient square_client.SquareClient
	cubeClient   cube_client.CubeClient
}

func (s *SumService) GetSum(a, b float32) float32 {
	squaredA := s.squareClient.GetSquare(a)
	cubedB := s.cubeClient.GetCube(b)
	return squaredA + cubedB
}

func NewSumService(
	squareClient square_client.SquareClient,
	cubeClient cube_client.CubeClient,
) *SumService {
	return &SumService{cubeClient: cubeClient, squareClient: squareClient}
}
