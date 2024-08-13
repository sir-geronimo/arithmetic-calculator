package usecases

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sir-geronimo/arithmetic-calculator/internal/domain/entities"
	"gorm.io/gorm"
)

type LoginUseCase struct {
	db *gorm.DB
}

type Token struct {
	AccessToken string `json:"access_token"`
}

func NewLoginUseCase(db *gorm.DB) *LoginUseCase {
	return &LoginUseCase{db}
}

func (u *LoginUseCase) Execute(username, password string) (*Token, error) {
	if username == "" || password == "" {
		return nil, ErrInvalidCredentials
	}

	// Retrieve user that matches `username`
	var user *entities.User
	err := u.db.
		Where("username = ?", username).
		First(&user).
		Error
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInvalidCredentials
	}

	if !user.VerifyPassword(password) {
		return nil, ErrInvalidCredentials
	}

	token, err := u.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &Token{AccessToken: token}, nil
}

func (u *LoginUseCase) generateToken(user *entities.User) (string, error) {
	exp, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		exp = 900
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":     os.Getenv("APP_HOST"),
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Second * time.Duration(exp)).Unix(),
		"user_id": user.ID,
	})

	secret := []byte(os.Getenv("JWT_SECRET"))

	return token.SignedString(secret)
}
