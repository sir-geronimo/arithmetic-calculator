package usecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/sir-geronimo/arithmetic-calculator/internal/domain/entities"
	"gorm.io/gorm"
)

type DeleteRecordUseCase struct {
	db *gorm.DB
}

func NewDeleteRecordUseCase(db *gorm.DB) *DeleteRecordUseCase {
	return &DeleteRecordUseCase{db}
}

func (u *DeleteRecordUseCase) Execute(recordID uuid.UUID) (*entities.Record, error) {
	var record entities.Record
	err := u.db.
		Where("id = ?", recordID).
		First(&record).
		Error
	if err != nil {
		return nil, ErrRecordNotFound
	}

	record.DeletedAt = time.Now()

	err = u.db.Save(record).Error
	if err != nil {
		return nil, ErrUnableToSaveRecord
	}

	return &record, nil
}
