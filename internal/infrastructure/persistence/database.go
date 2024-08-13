package persistence

import (
	"log"
	"os"
	"time"

	"github.com/sir-geronimo/arithmetic-calculator/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	db := GetConnection()

	_ = db.AutoMigrate(
		&entities.User{},
		&entities.Operation{},
		&entities.Record{},
	)

	// Uncomment for programmtic seeding instead
	// if db.Migrator().HasTable(&entities.User{}) {
	// 	err := db.First(&entities.User{}).Error
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		db.Create(&entities.User{
	// 			ID:       uuid.New(),
	// 			Username: "sir-geronimo",
	// 			Password: "Sup3rP4ssw0rd",
	// 			Status:   entities.UserActive,
	// 		})
	// 	}
	// }
}

// GetConnection returns a new database session
func GetConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Fatalf("Unable to connect to database. Error: %+v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 15)

	return db
}
