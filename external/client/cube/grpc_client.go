package client

type CubeGrpcClient struct{}

func (client *CubeGrpcClient) GetCube(num float32) float32 {
	return 1.0
}

func NewCubeGrpcClient() (*CubeGrpcClient, error) {
	return &CubeGrpcClient{}, nil
}
