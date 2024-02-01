package services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/brltd/delivery/domain"
	"github.com/brltd/delivery/dtos"
	"github.com/brltd/delivery/helpers"
	"github.com/brltd/delivery/logger"
	"github.com/brltd/delivery/repositories"
	"github.com/golang-jwt/jwt"
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

func (d DefaultUserService) AuthenticateUser() {

}

func generateToken(user domain.User) (*dtos.AuthResponse, error) {
	expiration := os.Getenv("EXP_HOUR")
	secret := os.Getenv("JWT_SECRET")

	exp, err := strconv.Atoi(expiration)

	if err != nil {
		return nil, err
	}

	expDate := time.Now().Add(time.Hour * time.Duration(exp))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"exp":   expDate.Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &dtos.AuthResponse{
		Token: tokenString,
		Exp:   expDate,
	}, nil
}
