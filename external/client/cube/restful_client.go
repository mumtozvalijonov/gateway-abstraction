package client

type CubeRestfulClient struct{}

func (client *CubeRestfulClient) GetCube(num float32) float32 {
	return 1.0
}

func NewCubeRestfulClient() (*CubeRestfulClient, error) {
	return &CubeRestfulClient{}, nil
}
