package service

import (
	"errors"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"time"
)


type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) Register(input dto.UserRegisterDto) (string, error) {
	
	hPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email: input.Email,
		Password: hPassword,
		Phone: input.Phone,
	})
	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
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

	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) isVerifiedUser(id uint) bool {

	user, err := s.Repo.FindUserById(id)

	return err == nil && user.Verified
}
func (s UserService) GetVerificationCode(e domain.User) (int, error) {

	if s.isVerifiedUser(e.ID) {
		return 0, errors.New("user is already verified")
	}

	code, err := s.Auth.GenerateCode()
	if err != nil {
		return 0, err
	}

	user := domain.User{
		Code: code,
		Expiry: time.Now().Add(time.Minute * 15),
	}

	_, err = s.Repo.UpdateUser(e.ID, user)
	if err != nil {
		return 0, errors.New("error while updating user")
	}

	// TODO: send code in SMS

	return code, nil
}

func (s UserService) VerifyCode(id uint, code int) error {

	if s.isVerifiedUser(id) {
		return errors.New("user is already verified")
	}

	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return errors.New("user does not exist")
	}

	if user.Code != code {
		return errors.New("invalid verification code")
	}
	if user.Expiry.Before(time.Now()) {
		return errors.New("verification code expired")
	}

	updateUser := domain.User{
		Verified: true,
	}
	_, err = s.Repo.UpdateUser(id, updateUser)
	if err != nil {
		return errors.New("error while verifying user")
	}

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