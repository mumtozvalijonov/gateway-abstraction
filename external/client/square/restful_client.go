package client

type SquareRestfulClient struct{}

func (client *SquareRestfulClient) GetSquare(num float32) float32 {
	return 1.0
}

func NewSquareRestfulClient() (*SquareRestfulClient, error) {
	return &SquareRestfulClient{}, nil
}
