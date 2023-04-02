package utils

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	""
)

const (
	host     = "localhost"
	port     = 5432
	user    = "postgres"
	password = "postgres"
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
}

