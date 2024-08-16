package entities

import (
	"time"

	"github.com/google/uuid"
)

// Record represents the result of an Operation.
type Record struct {
	ID                uuid.UUID `json:"id" gorm:"type:uuid; primaryKey"`
	OperationID       uuid.UUID `json:"operation_id" gorm:"not null;"`
	UserID            uuid.UUID `json:"user_id" gorm:"not null;"`
	Amount            int       `json:"amount" gorm:"not null;"`
	UserBalance       int       `json:"user_balance" gorm:"not null;"`
	OperationResponse string    `json:"operation_response" gorm:"not null;"`
	Date              time.Time `json:"date" gorm:"not null; default:NOW();"`
	DeletedAt         time.Time `json:"deleted_at" gorm:"default:null"`

	Operation *Operation `json:"operation" gorm:"not null;"`
	User      *User      `json:"-" gorm:"not null;"`
}

func NewRecord(
	operationID, userID uuid.UUID,
	amount int,
	userBalance int,
) *Record {
	return &Record{
		ID:          uuid.New(),
		OperationID: operationID,
		UserID:      userID,
		Amount:      amount,
		UserBalance: userBalance,
		Date:        time.Now(),
	}
}
