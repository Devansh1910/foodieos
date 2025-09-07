package getOutletFood

import (
	"log"
	"time"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/datatypes"
)

var DB *gorm.DB

// OutletFood is a simple table to store full Response JSON per outlet
type OutletFood struct {
	ID        uint           `gorm:"primaryKey"`
	OutletID  int            `gorm:"index;not null"`
	Data      datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// InitDB connects to Postgres and migrates the OutletFood table.
// It reads DB config from environment variable DATABASE_DSN if present,
// otherwise uses a sensible local default — change for production.
func InitDB() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Println("DATABASE_DSN not set, skipping DB init (using mock data only)")
		return
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("failed to connect database:", err)
		DB = nil
		return
	}

	if err := DB.AutoMigrate(&OutletFood{}); err != nil {
		log.Println("auto migrate failed:", err)
		DB = nil
		return
	}

	log.Println("database connected & migrated")
}

