package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBmap = initPG()

func initPG() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbmap, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	checkErr(err, "failed to connect database")

	dbmap.AutoMigrate(User{})

	return dbmap
}

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "db_teste"
)

func checkErr(err error, msg string) *ServerError {
	if err != nil {
		log.Fatalln(msg, err)
		return &ServerError{
			Message: "Internal Server Error",
			Error:   err,
		}
	}
	return nil
}
