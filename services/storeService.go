package services

type StoreService interface {
}

type DefaultStoreService struct {
}

func NewStoreService() DefaultStoreService {
	return DefaultStoreService{}
}
