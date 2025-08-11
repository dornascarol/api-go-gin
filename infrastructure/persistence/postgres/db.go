package postgres

import (
	"log"

	"github.com/dornascarol/api-go-gin/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() *gorm.DB {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error connecting to database")
	}
	db.AutoMigrate(&entities.Singer{})
	return db
}
