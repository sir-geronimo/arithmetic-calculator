package entities

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserStatus uint8

const (
	UserInactive UserStatus = 0
	UserActive   UserStatus = 1
)

var (
	ErrUnableToCreateUser = errors.New("unable to create user")
)

// User represents the main actor that perform operations.
type User struct {
	ID       uuid.UUID  `json:"id" gorm:"type:uuid; primaryKey;"`
	Username string     `json:"username" gorm:"type:varchar(50); unique; not null;"`
	Password string     `json:"-" gorm:"not null;"`
	Status   UserStatus `json:"status" gorm:"not null;"`
}

func NewUser(
	id uuid.UUID,
	username, password string,
	status UserStatus,
) *User {
	return &User{
		ID:       id,
		Username: username,
		Password: password,
		Status:   status,
	}
}

// VerifyPassword checks if the provided password matches the user encrypted password
func (u *User) VerifyPassword(password string) (valid bool) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

func (u *User) BeforeCreate(*gorm.DB) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln("Err: Unable to create user", err.Error())
		return err
	}

	u.Password = string(bytes)

	return nil
}
