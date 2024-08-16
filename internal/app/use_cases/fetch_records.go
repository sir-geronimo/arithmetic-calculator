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

type FetchRecordsResult struct {
	Data  []*entities.Record `json:"data"`
	Page  int                `json:"page"`
	Total int64              `json:"total"`
}

func (u *FetchRecordsUseCase) Execute(userID uuid.UUID, opts *FetchRecordsOptions) (*FetchRecordsResult, error) {
	var records []*entities.Record
	q := u.db.Where("user_id = ? AND deleted_at IS NULL", userID)
	filter := "%" + opts.Filter + "%"

	if opts.Filter != "" {
		q.Where("amount::varchar ILIKE ? OR user_balance::varchar ILIKE ? OR operation_response ILIKE ?", filter, filter, filter)
	}

	offset := (opts.Page - 1) * opts.PerPage
	err := q.
		Order(clause.OrderByColumn{Column: clause.Column{Name: "date"}, Desc: !opts.OrderAsc}).
		Offset(offset).
		Limit(opts.PerPage).
		Preload("Operation").
		Find(&records).
		Error
	if err != nil {
		return nil, ErrUnableToFetchRecords
	}

	var total int64
	q = u.db.
		Model(&entities.Record{}).
		Where("user_id = ? AND deleted_at IS NULL", userID)
	if opts.Filter != "" {
		q.Where("amount::varchar ILIKE ? OR user_balance::varchar ILIKE ? OR operation_response ILIKE ?", filter, filter, filter)
	}

	err = q.Count(&total).Error
	if err != nil {
		return nil, ErrUnableToFetchRecords
	}

	result := &FetchRecordsResult{
		Data:  records,
		Page:  opts.Page,
		Total: total,
	}

	return result, nil
}
