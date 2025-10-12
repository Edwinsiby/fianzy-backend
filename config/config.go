package config

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("fianzy.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect database: ", err)
	}

	fmt.Println("✅ Database connected successfully!")
}
