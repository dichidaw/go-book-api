package handlers

import (
	"fmt"
	dm "go-book-api/datamodels"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable dbname=postgres",
		dbHost, dbUser, dbPassword, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the 'postgres' database: ", err)
	}

	createOrUseDB(db, dbName)
	CloseDBConn(db)

	newDB := changeDB(dbHost, dbUser, dbPassword, dbPort, dbName)
	if os.Getenv("RUN_MIGRATIONS") == "TRUE" {
		migrateDatabase(newDB)
	}
	return newDB
}

func createOrUseDB(db *gorm.DB, dbName string) {
	var dbExists bool
	db.Raw("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = ?)", dbName).Scan(&dbExists)
	if !dbExists {
		createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", dbName)
		if err := db.Exec(createDatabaseCommand).Error; err != nil {
			log.Fatal("Failed to create database: ", err)
		}
		log.Infof("Database %s created successfully!", dbName)
	}
}

func CloseDBConn(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB from GORM instance: ", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatal("Failed to close DB connection: ", err)
	}
}

func changeDB(dbHost string, dbUser string, dbPassword string, dbPort string, dbName string) *gorm.DB {
	newDSN := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable dbname=%s",
		dbHost, dbUser, dbPassword, dbPort, dbName)

	newDB, err := gorm.Open(postgres.Open(newDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the new database: ", err)
	}

	log.Infof("Successfully connected to the database %s", dbName)
	return newDB
}

func migrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&dm.User{}, &dm.Book{}, &dm.Borrowing{}, &dm.Writer{}, &dm.BookWriter{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	log.Info("Database migrated successfully!")
}
