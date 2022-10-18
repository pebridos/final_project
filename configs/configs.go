package configs

import (
	"fmt"
	"log"

	"final_project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "P@ssw0rd"
	dbname   = "my_gram"
	db       *gorm.DB
	err      error
)

func StartDB() {
	psqInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(psqInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to database because: ", err)
	}
	db.Debug().AutoMigrate(models.User{}, models.SocialMedia{}, models.Photo{}, models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
