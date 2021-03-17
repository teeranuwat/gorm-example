package main

import (
	"fmt"
	"gorm-example/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// do run migrate
	err := godotenv.Load("/Users/teeranuwatueaprasert/Me/workspace/gorm-example/config/.env")
	if err != nil {
		log.Panic(err.Error())
	}
	customer := model.Customer{}
	card := model.Card{}

	pool, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("DBHOSTNAME"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPASSWORD"), os.Getenv("SSLMODE"))), &gorm.Config{})

	if err != nil {
		log.Println(err.Error())
	}

	err = pool.AutoMigrate(&customer, &card)
	if err != nil {
		log.Println(err.Error())
	}
}
