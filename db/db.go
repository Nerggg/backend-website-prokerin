package db

import (
	"backend-prokerin/config"
	"backend-prokerin/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// get from config
	host := config.Config.Database.Host
	username := config.Config.Database.Username
	password := config.Config.Database.Password
	databaseName := config.Config.Database.Db
	port := config.Config.Database.Port

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, databaseName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		//NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		//Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func init() {
	db := Connect()
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.UserAccount{})
	db.AutoMigrate(&models.Proker{})
}
func Close(db *gorm.DB) {
	sqlDB, _ := db.DB()

	// Close
	sqlDB.Close()
}
