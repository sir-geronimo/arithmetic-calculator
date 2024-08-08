package entities

import (
	"time"

	"github.com/google/uuid"
)

// Record represents the result of an Operation.
type Record struct {
	ID                uuid.UUID `json:"id"`
	OperationID       string    `json:"operation_id"`
	UserID            string    `json:"user_id"`
	Amount            int       `json:"amount"`
	UserBalance       int       `json:"user_balance"`
	OperationResponse string    `json:"operation_response"`
	Date              time.Time `json:"date"`
	CreatedAt         time.Time `json:"created_at"`
	DeletedAt         time.Time `json:"deleted_at"`
}
