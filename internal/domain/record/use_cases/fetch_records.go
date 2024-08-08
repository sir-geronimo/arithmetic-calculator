package usecases

import (
	"github.com/sir-geronimo/arithmetic-calculator/internal/domain/record/entities"
	"gorm.io/gorm"
)

type FetchRecordsUseCase struct {
	db *gorm.DB
}

func NewFetchRecordsUseCase(db *gorm.DB) *FetchRecordsUseCase {
	return &FetchRecordsUseCase{db}
}

type FetchRecordsOptions struct {
	Page    int
	PerPage int
	Filter  string
}

func (u *FetchRecordsUseCase) Execute(userID string, options *FetchRecordsOptions) ([]entities.Record, error) {
	var records []entities.Record
	err := u.db.
		Debug().
		Where("user_id = ? AND deleted_at IS NULL", userID).
		Where("amount LIKE ? OR user_balance LIKE ? OR operation_response LIKE ?", options.Filter, options.Filter, options.Filter).
		Order("created_at DESC").
		Offset(options.Page).
		Limit(options.PerPage).
		Find(&records).
		Error
	if err != nil {
		return nil, ErrUnableToFetchRecords
	}

	// var result []entities.Record
	// for _, v := range records {
	// 	if strings.Contains(strconv.Itoa(int(v.Amount)), options.Filter) {
	// 		result = append(result, v)
	// 	} else if strings.Contains(strconv.Itoa(int(v.UserBalance)), options.Filter) {
	// 		result = append(result, v)
	// 	} else if strings.Contains(v.OperationResponse, options.Filter) {
	// 		result = append(result, v)
	// 	}
	// }

	return records, nil
}
