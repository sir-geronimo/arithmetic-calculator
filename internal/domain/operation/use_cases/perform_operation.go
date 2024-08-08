package usecases

import (
	"github.com/sir-geronimo/arithmetic-calculator/internal/domain/operation/entities"
	"gorm.io/gorm"
)

type PerformOperationUseCase struct {
	db *gorm.DB
}

func NewPerformOperationUseCase(db *gorm.DB) *PerformOperationUseCase {
	return &PerformOperationUseCase{db}
}

func (u *PerformOperationUseCase) Execute(name entities.OperationName) (*entities.Operation, error) {
	if !name.IsValid() {
		return nil, ErrInvalidOperation
	}

	switch name {
	case entities.OperationAddition:
		break
	case entities.OperationSubtraction:
		break
	case entities.OperationMultiplication:
		break
	case entities.OperationDivision:
		break
	case entities.OperationSquareRoot:
		break
	case entities.OperationRandomString:
		break
	}

	return nil, nil
}
