package usecases

import (
	"github.com/google/uuid"
	"github.com/sir-geronimo/arithmetic-calculator/internal/domain/entities"
	"gorm.io/gorm"
)

type CreateOperationUseCase struct {
	db *gorm.DB
}

func NewCreateOperationUseCase(db *gorm.DB) *CreateOperationUseCase {
	return &CreateOperationUseCase{db}
}

func (u *CreateOperationUseCase) Execute(userID uuid.UUID, name string) (*entities.Operation, error) {
	n := entities.OperationName(name)
	if !n.IsValid() {
		return nil, ErrInvalidOperationType
	}

	operation := entities.NewOperation(n)
	_, err := computeBalance(u.db, userID, operation)
	if err != nil {
		return nil, err
	}

	err = u.db.Save(&operation).Error
	if err != nil {
		return nil, ErrUnableToSaveOperation
	}

	return operation, nil
}

// computeBalance returns the new user balance based on the last user balance available.
//
// If balance is not enough to cover operation, returns ErrInsufficientBalance.
func computeBalance(
	db *gorm.DB,
	userID uuid.UUID,
	operation *entities.Operation,
) (int, error) {
	// Retrieve last user record
	var record entities.Record
	err := db.
		Where("user_id = ?", userID).
		Limit(1).
		Order("date DESC").
		Find(&record).
		Error
	if err != nil {
		return 0, ErrUnableToFindRecord
	}

	var balance int
	if record.ID == uuid.Nil {
		balance = 100 - operation.Cost
	} else {
		balance = record.UserBalance - operation.Cost
	}

	if balance < 0 {
		return 0, ErrInsufficientBalance
	}

	return balance, nil
}
