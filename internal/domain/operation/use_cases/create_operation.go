package usecases

import (
	"github.com/sir-geronimo/arithmetic-calculator/internal/domain/operation/entities"
	"gorm.io/gorm"
)

type CreateOperationUseCase struct {
	db *gorm.DB
}

func NewCreateOperationUseCase(db *gorm.DB) *CreateOperationUseCase {
	return &CreateOperationUseCase{db}
}

func (u *CreateOperationUseCase) Execute(name string) (*entities.Operation, error) {
	n := entities.OperationName(name)
	if !n.IsValid() {
		return nil, ErrInvalidOperation
	}

	operation := entities.NewOperation(n)

	err := u.db.Save(&operation).Error
	if err != nil {
		return nil, ErrUnableToSaveOperation
	}

	return operation, nil
}
