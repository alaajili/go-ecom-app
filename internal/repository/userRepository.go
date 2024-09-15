package repository

import (
	"errors"
	"go-ecommerce-app/internal/domain"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)



type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, user domain.User) (domain.User, error)

	CreateBankAccount(bank_account domain.BankAccount) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateBankAccount(bank_account domain.BankAccount) error {
	return r.db.Create(&bank_account).Error
}

func (r userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return domain.User{}, errors.New("error creating user")
	}

	return user, nil
}

func (r userRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "email=?", email).Error
	if err != nil {
		log.Printf("Error finding user by email: %v", err)
		return domain.User{}, errors.New("error finding user by email")
	}

	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, id).Error
	if err != nil {
		log.Printf("Error finding user by id: %v", err)
		return domain.User{}, errors.New("error finding user by id")
	}

	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return domain.User{}, errors.New("updating user failed")
	}

	return user, nil
}

