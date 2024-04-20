package mysql

import (
	"context"
	"fmt"
	"os"

	"github.com/velicanercan/simple-user-mgmt/domain"
	"github.com/velicanercan/simple-user-mgmt/logger"
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
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		logger.Log("Failed to migrate the schema!", "panic", err.Error())
	}
	return &MySQLDatabase{DB: db}
}

// Close closes the MySQL database connection
func (db *MySQLDatabase) Close(ctx context.Context) {
	sqlDB, err := db.DB.DB()
	if err != nil {
		logger.Log("error", "Failed to close the database connection!", err.Error())
	}
	sqlDB.Close()
}

// InsertUser inserts a user into the MySQL database
func (db *MySQLDatabase) InsertUser(ctx context.Context, user *domain.User) error {
	result := db.DB.WithContext(ctx).Create(user)
	return result.Error
}

// GetAllUsers retrieves all users from the MySQL database
func (db *MySQLDatabase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	result := db.DB.WithContext(ctx).Find(&users)
	return users, result.Error
}

// GetUser retrieves a user from the MySQL database by ID
func (db *MySQLDatabase) GetUser(ctx context.Context, id int) (*domain.User, error) {
	user := domain.User{}
	result := db.DB.WithContext(ctx).First(&user, id)
	return &user, result.Error
}

// UpdateUser updates a user in the MySQL database
func (db *MySQLDatabase) UpdateUser(ctx context.Context, id int, user *domain.User) error {
	result := db.DB.Model(&domain.User{}).WithContext(ctx).Where("id = ?", id).Updates(user)
	return result.Error
}

// DeleteUser deletes a user from the MySQL database by ID
func (db *MySQLDatabase) DeleteUser(ctx context.Context, id int) error {
	result := db.DB.WithContext(ctx).Delete(&domain.User{}, id)
	return result.Error
}
