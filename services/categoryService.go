package services

type CategoryService interface {
}

type DefaultCategoryService struct {
}

func NewCategoryService() DefaultCategoryService {
	return DefaultCategoryService{}
}
