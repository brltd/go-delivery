package services

import (
	"fmt"
	"time"

	"github.com/brltd/delivery/domain"
	"github.com/brltd/delivery/dtos"
	"github.com/brltd/delivery/helpers"
	"github.com/brltd/delivery/logger"
	"github.com/brltd/delivery/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(req dtos.CreateUserRequest) (*dtos.CreateUserResponse, *helpers.ApiError)
}

type DefaultUserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) DefaultUserService {
	return DefaultUserService{
		userRepository: userRepository,
	}
}

func (d DefaultUserService) CreateUser(req dtos.CreateUserRequest) (*dtos.CreateUserResponse, *helpers.ApiError) {
	err := req.Validate()

	if err != nil {
		return nil, err
	}

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if errHash != nil {
		logger.Error(fmt.Sprintf("Error while hashing the password: %+v", errHash.Error()))
		return nil, helpers.InternalServerError("Unexpected error")
	}

	newUser, err := d.userRepository.CreateUser(domain.User{
		Name:        req.Name,
		Email:       req.Email,
		Hash:        string(hashedPassword),
		DateCreated: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	response := newUser.ToCreateUserResponseDto()

	return &response, nil
}
