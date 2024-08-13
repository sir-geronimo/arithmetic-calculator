package usecases

import (
	"github.com/google/uuid"
	"github.com/sir-geronimo/arithmetic-calculator/internal/domain/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FetchRecordsUseCase struct {
	db *gorm.DB
}

func NewFetchRecordsUseCase(db *gorm.DB) *FetchRecordsUseCase {
	return &FetchRecordsUseCase{db}
}

type FetchRecordsOptions struct {
	Page     int
	PerPage  int
	OrderAsc bool
	Filter   string
}

func (u *FetchRecordsUseCase) Execute(userID uuid.UUID, opts *FetchRecordsOptions) ([]*entities.Record, error) {
	var records []*entities.Record
	q := u.db.Where("user_id = ? AND deleted_at IS NULL", userID)

	if opts.Filter != "" {
		q.Where("amount = ? OR user_balance = ? OR operation_response LIKE ?", opts.Filter, opts.Filter, opts.Filter)
	}

	err := q.
		Order(clause.OrderByColumn{Column: clause.Column{Name: "date"}, Desc: !opts.OrderAsc}).
		Offset(opts.Page - 1).
		Limit(opts.PerPage).
		Find(&records).
		Error
	if err != nil {
		return nil, ErrUnableToFetchRecords
	}

	return records, nil
}
