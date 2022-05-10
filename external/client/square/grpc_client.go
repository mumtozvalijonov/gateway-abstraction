package client

type SquareGrpcClient struct{}

func (client *SquareGrpcClient) GetSquare(num float32) float32 {
	return 1.0
}

func NewSquareGrpcClient() (*SquareGrpcClient, error) {
	return &SquareGrpcClient{}, nil
}
