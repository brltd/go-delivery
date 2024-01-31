package services

type ProductService interface {
}

type DefaultProductService struct {
}

func NewProductService() DefaultProductService {
	return DefaultProductService{}
}
