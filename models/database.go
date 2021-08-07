package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=bdacc password=132596 sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Acc{})

	DB = db
}
