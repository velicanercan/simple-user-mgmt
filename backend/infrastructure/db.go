package infrastructure

import (
	"fmt"
	"os"

	"github.com/velicanercan/simple-user-mgmt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	MySQLDB *gorm.DB
}

// NewMySQLDatabase intializes and returns mysql db
func NewMySQLDatabase() Database {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")
	}
	
	return Database{
		MySQLDB: db,
	}
}

func (db *Database) Close() {
	sqlDB, err := db.MySQLDB.DB()
	if err != nil {
		panic("Failed to close database connection!")
	}
	sqlDB.Close()
}

func (db *Database) Migrate() {
	db.MySQLDB.AutoMigrate(&models.User{})
}
