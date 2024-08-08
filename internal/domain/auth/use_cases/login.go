package usecases

import (
	"github.com/sir-geronimo/arithmetic-calculator/internal/domain/auth/entities"
	"gorm.io/gorm"
)

type LoginUseCase struct {
	db *gorm.DB
}

func NewLoginUseCase(db *gorm.DB) *LoginUseCase {
	return &LoginUseCase{db}
}

func (u *LoginUseCase) Execute(username, password string) (*entities.User, error) {
	var user *entities.User

	err := u.db.
		Where("username = ?", username).
		First(user).
		Error
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !user.VerifyPassword(password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
