package usecases

import (
	"fmt"
	"math"
	"strconv"

	"github.com/google/uuid"
	"github.com/sir-geronimo/arithmetic-calculator/internal/domain/entities"
	"github.com/sir-geronimo/arithmetic-calculator/internal/infrastructure/external"
	"gorm.io/gorm"
)

type PerformOperationUseCase struct {
	db     *gorm.DB
	strGen external.StringGenerator
}

type PerformOperationRequest struct {
	OperationID uuid.UUID
	UserID      uuid.UUID
	Num1        int
	Num2        int
}

func NewPerformOperationUseCase(db *gorm.DB) *PerformOperationUseCase {
	return &PerformOperationUseCase{
		db:     db,
		strGen: external.NewHTTPStringGenerator(),
	}
}

func (u *PerformOperationUseCase) Execute(req *PerformOperationRequest) (record *entities.Record, err error) {
	var operation *entities.Operation

	// Validate requested operation exists
	err = u.db.
		Where("id = ?", req.OperationID).
		Preload("Records", "deleted_at IS NULL").
		First(&operation).
		Error
	if err != nil {
		return nil, ErrOperationNotFound
	}

	if operation.IsPerformed() {
		return nil, ErrOperationAlreadyPerformed
	}

	balance, err := computeBalance(u.db, req.UserID, operation)
	if err != nil {
		return
	}

	// Perform operation
	record = entities.NewRecord(
		req.OperationID,
		req.UserID,
		operation.Cost,
		balance,
	)

	var operationResponse string
	switch operation.Name {
	case entities.OperationAddition:
		operationResponse, err = u.add(req)
	case entities.OperationSubtraction:
		operationResponse, err = u.subtract(req)
	case entities.OperationMultiplication:
		operationResponse, err = u.multiply(req)
	case entities.OperationDivision:
		operationResponse, err = u.divide(req)
	case entities.OperationSquareRoot:
		operationResponse, err = u.sqrt(req)
	case entities.OperationRandomString:
		operationResponse, err = u.randomStr()
	}

	if err != nil {
		return
	}

	record.OperationResponse = operationResponse

	// Save record to database
	u.db.Transaction(func(tx *gorm.DB) error {
		err = u.db.
			Save(&record).
			Error
		if err != nil {
			return ErrUnableToPerformOperation
		}

		// Return record
		err = u.db.
			Preload("Operation").
			First(&record).
			Error
		if err != nil {
			return ErrUnableToPerformOperation
		}

		return nil
	})

	return record, err
}

func (u *PerformOperationUseCase) add(req *PerformOperationRequest) (string, error) {
	if req.Num1 == 0 && req.Num2 == 0 {
		return "", fmt.Errorf("%w: both number must not be `0`", ErrInvalidOperationPayload)
	}

	res := req.Num1 + req.Num2

	return strconv.Itoa(res), nil
}

func (u *PerformOperationUseCase) subtract(req *PerformOperationRequest) (string, error) {
	if req.Num1 == 0 && req.Num2 == 0 {
		return "", fmt.Errorf("%w: both number must not be `0`", ErrInvalidOperationPayload)
	}

	res := req.Num1 - req.Num2

	return strconv.Itoa(res), nil
}

func (u *PerformOperationUseCase) multiply(req *PerformOperationRequest) (string, error) {
	if req.Num1 == 0 && req.Num2 == 0 {
		return "", fmt.Errorf("%w: both number must not be `0`", ErrInvalidOperationPayload)
	}

	res := req.Num1 * req.Num2

	return strconv.Itoa(res), nil
}

func (u *PerformOperationUseCase) divide(req *PerformOperationRequest) (string, error) {
	if req.Num1 == 0 && req.Num2 == 0 {
		return "", fmt.Errorf("%w: both number must not be `0`", ErrInvalidOperationPayload)
	}
	if req.Num2 == 0 {
		return "", fmt.Errorf("%w: invalid divisor. Divisor must not be `0`", ErrInvalidOperationPayload)
	}

	res := req.Num1 / req.Num2

	return strconv.Itoa(res), nil
}

func (u *PerformOperationUseCase) sqrt(req *PerformOperationRequest) (string, error) {
	if req.Num1 == 0 {
		return "", fmt.Errorf("%w: number must not be `0`", ErrInvalidOperationPayload)
	}
	if req.Num1 < 0 {
		return "", fmt.Errorf("%w: number must not be negative", ErrInvalidOperationPayload)
	}

	res := int(math.Sqrt(float64(req.Num1)))

	return strconv.Itoa(res), nil
}

func (u *PerformOperationUseCase) randomStr() (string, error) {
	str, err := u.strGen.Generate(&external.GenerateStringOptions{
		Len:    12,
		Unique: true,
	})
	if err != nil {
		return "", err
	}

	return str, nil
}
