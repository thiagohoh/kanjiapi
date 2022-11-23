package database

import (
	"github.com/kyary/kanjiapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	conn := "host=localhost user=postgres password= dbname= port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(conn))

	if err != nil {
		panic(err.Error())
	}

	DB.AutoMigrate(&models.Kanji{})
	DB.AutoMigrate(&models.Jlptlvl{})
}
