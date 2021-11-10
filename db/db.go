package db

import (
	"log"
	"music/app/model"
	"os"

	"github.com/jinzhu/gorm"
)

func Init() *gorm.DB {
	dsn := os.Getenv("DSN")
	log.Println("Connecting to database...")
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Println("Could not connect to database")
		log.Fatal(err)
	}
	log.Println("Connected to database")
	log.Println("Migrating database...")
	mdb := migrate(db)
	log.Println("Migrated database")
	return mdb
}

func migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.Track{})
	return db
}
