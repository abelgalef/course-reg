package storage

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/abelgalef/course-reg/pkg/models"
)

type MainDB struct {
	Connection *gorm.DB
	Type       string
}

func NewDatabaseConnection(createMySqlDB bool) *MainDB {
	var db *gorm.DB
	var err error

	// SUPPORT FOR MULTIPLE TYPES OF DATABASES
	if !createMySqlDB {
		// HARD CODED CREDENTIALS
		dsn := "courseReg:password@tcp(localhost:3306)/courseReg?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("\nStorage: Storage.go: Mysql connection error: %v", err)
		}
	} else {
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("\nStorage: Storage.go: SQLite connection error: %v", err)
		}
	}

	// COLLECT ALL MODELS HERE
	// THE ORDER OF TABLES MATTERS HERE BECAUSE OF FOREGIN KEY DEPENDENCIES
	tables := []interface{}{
		&models.Right{},
		&models.Course{},
		&models.Files{},
		&models.User{},
		&models.Department{},
		&models.Role{},
	}

	if err1 := db.AutoMigrate(tables...); err != nil {
		log.Fatalf("\nStorage: Storage.go: Could not create tables: %v", err1)
	}

	fmt.Println("Initializing the database with permissions and superusers")
	preloadDatabase(db)

	if createMySqlDB {
		return &MainDB{db, "MYSQL"}
	} else {
		return &MainDB{db, "SQLITE"}
	}
}

func preloadDatabase(db *gorm.DB) {
	var role models.Role

	// CHECK IF THE DATABASE HAS BEEN INITIATED BEFORE BY TRYING TO GET ANY USER
	if err := db.Take(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // IF THIS IS TRUE THEN THE DATABASE IS EMPTY
			if err := db.Create(&models.NeededRole).Error; err != nil {
				log.Fatalf("\nStorage: Storage.go: Could not create the roles and permissions needed: %v", err)
			}

			fmt.Println("Database preloaded")
		} else {
			log.Fatalf("\nStorage: Storage.go: %v", err)
		}
	}
}
