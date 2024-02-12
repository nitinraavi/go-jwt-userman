package intializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DBURI")), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		panic("Failed to Connect to DB")
	}

}
