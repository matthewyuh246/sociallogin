package repository

import (
	"github.com/matthewyuh246/socallogin/internal/domain"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *domain.User, email string) error
	CreateUser(user *domain.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRespository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetUserByEmail(user *domain.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *domain.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
