package usecases

import (
	"github.com/google/uuid"
	"github.com/sir-geronimo/arithmetic-calculator/internal/domain/entities"
	"gorm.io/gorm"
)

const (
	InitialBalance int = 100
)

type GetBalanceUseCase struct {
	db *gorm.DB
}

func NewGetBalanceUseCase(db *gorm.DB) *GetBalanceUseCase {
	return &GetBalanceUseCase{db}
}

func (u *GetBalanceUseCase) Execute(userID uuid.UUID) (balance int, err error) {
	balance = InitialBalance
	record, err := retrieveLastRecord(u.db, userID)
	if err != nil {
		return 0, ErrUnableToGetBalance
	}

	if record.ID != uuid.Nil {
		balance = record.UserBalance
	}

	return
}

func retrieveLastRecord(db *gorm.DB, userID uuid.UUID) (record entities.Record, err error) {
	err = db.
		Where("user_id = ?", userID).
		Limit(1).
		Order("date DESC").
		Find(&record).
		Error

	return
}
