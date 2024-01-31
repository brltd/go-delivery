package services

type StoreStatusService interface {
}

type DefaultStoreStatusService struct {
}

func NewStoreStatusService() DefaultStoreStatusService {
	return DefaultStoreStatusService{}
}
