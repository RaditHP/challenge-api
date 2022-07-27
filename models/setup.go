package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupDB() *gorm.DB {
	USER := "postgres"
	PASS := "docker"
	HOST := "localhost"
	PORT := 5432
	DBNAME := "challengedb"
	URL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", HOST, USER, PASS, DBNAME, PORT)
	db, err := gorm.Open("postgres", URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}
