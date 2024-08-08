package entities

import (
	"math/rand"

	"github.com/google/uuid"
)

// Operation represent an arithmetic calculation performed by the users.
type Operation struct {
	ID   uuid.UUID     `json:"id" gorm:"type:uuid; primarykey;"`
	Name OperationName `json:"name" gorm:"type:varchar(50);"`
	Cost int           `json:"cost"`
}

type OperationName string

const (
	OperationAddition       OperationName = "addition"
	OperationSubtraction    OperationName = "subtraction"
	OperationMultiplication OperationName = "multiplication"
	OperationDivision       OperationName = "division"
	OperationSquareRoot     OperationName = "square_root"
	OperationRandomString   OperationName = "random_string"
)

func (n OperationName) IsValid() bool {
	switch n {
	case
		OperationAddition,
		OperationSubtraction,
		OperationMultiplication,
		OperationDivision,
		OperationSquareRoot,
		OperationRandomString:
		return true
	default:
		return false
	}
}

func NewOperation(name OperationName) *Operation {
	return &Operation{
		ID:   uuid.New(),
		Name: name,
		Cost: rand.Intn(50),
	}
}
