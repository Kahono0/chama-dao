package utils

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Kahono0/chama-dao/models"
)

const (
	host     = "localhost"
	port     = 5432
	user    = "root"
	password = "root"
	dbname   = "chama"
)

var DB *gorm.DB

func ConnectDB() {
	dbStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dbStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	db.AutoMigrate(&models.Transaction{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Proposal{})
	db.AutoMigrate(&models.Organization{})

}

