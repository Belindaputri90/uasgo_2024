package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"uasgo_2024/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dataSourceName := "root:@tcp(127.0.0.1:3306)/uasgo_2024?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate the tables
	migrateDatabase()

	fmt.Println("Connected to database")
}

func migrateDatabase() {
	// Auto-migrate the tables
	err := DB.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		panic("failed to auto-migrate tables")
	}
}

// GetUserByUsername mengambil data user berdasarkan username
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
