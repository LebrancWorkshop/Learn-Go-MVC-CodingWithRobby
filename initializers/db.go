package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")
	dial := postgres.Open(dsn)
	DB, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = DB

}