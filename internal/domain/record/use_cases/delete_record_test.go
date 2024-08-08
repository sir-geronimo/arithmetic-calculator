package usecases_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	usecases "github.com/sir-geronimo/arithmetic-calculator/internal/domain/record/use_cases"
	"gorm.io/gorm"
)

func TestDeleteRecordUseCase(t *testing.T) {
	recordID := uuid.New()

	var db *gorm.DB
	u := usecases.NewDeleteRecordUseCase(db)

	t.Run("should soft-delete the `Record`", func(t *testing.T) {
		expected := time.Now()

		actual, _ := u.Execute(recordID)

		if actual.DeletedAt != expected {
			t.Errorf("got %s, want %s", actual.DeletedAt.String(), expected.String())
		}
	})

	t.Run("should fail if `Record` does not belong to `User`", func(t *testing.T) {
		// TODO: Write test
	})
}
