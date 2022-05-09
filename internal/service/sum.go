package service

type SumService struct {
}

func (s SumService) GetSum(a, b int32) int32 {
	return a + b
}

func NewSumService() *SumService {
	return &SumService{}
}
