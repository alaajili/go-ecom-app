package service

import (
	"errors"
	"fmt"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"log"
)


type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) Register(input dto.UserRegisterDto) (string, error) {
	log.Println(input)

	user, err := s.Repo.CreateUser(domain.User{
		Email: input.Email,
		Password: input.Password,
		Phone: input.Phone,
	})

	// generate token
	log.Println(user)
	userInfo := fmt.Sprintf("%d|%s|%s", user.ID, user.Email, user.Phone)

	return userInfo, err
}

func (s UserService) FindUserByEmail(email string) (*domain.User, error) {
	
	user, err := s.Repo.FindUserByEmail(email)

	return &user, err
}

func (s UserService) Login(email, password string) (string, error) {
	
	user, err := s.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email")
	}

	return user.Email, nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code string) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]any, error) {
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]any, error) {
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]any, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uId uint) ([]any, error) {
	return nil, nil
}