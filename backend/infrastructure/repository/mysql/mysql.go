package mysql

import (
	"fmt"
	"os"

	"github.com/velicanercan/simple-user-mgmt/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDatabase struct holds the MySQL database connection
type MySQLDatabase struct {
	DB *gorm.DB
}

// NewMySQLDatabase intializes and returns mysql db
func NewMySQLDatabase() *MySQLDatabase {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")
	}
	// migrate the schema
	db.AutoMigrate(&domain.User{})
	return &MySQLDatabase{DB: db}
}

// Close closes the MySQL database connection
func (db *MySQLDatabase) Close() {
	sqlDB, err := db.DB.DB()
	if err != nil {
		panic("Failed to close database connection!")
	}
	sqlDB.Close()
}

// InsertUser inserts a user into the MySQL database
func (db *MySQLDatabase) InsertUser(user *domain.User) error {
	result := db.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllUsers retrieves all users from the MySQL database
func (db *MySQLDatabase) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

// GetUser retrieves a user from the MySQL database by ID
func (db *MySQLDatabase) GetUser(id int) (*domain.User, error) {
	user := domain.User{}
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return &user, result.Error
	}
	return &user, nil
}

// UpdateUser updates a user in the MySQL database
func (db *MySQLDatabase) UpdateUser(id int, user *domain.User) error {
	result := db.DB.Model(&domain.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser deletes a user from the MySQL database by ID
func (db *MySQLDatabase) DeleteUser(id int) error {
	result := db.DB.Delete(&domain.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
