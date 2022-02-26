package utils

import (
	"fmt"
	"group-project2/configs"
	I "group-project2/entities/image"
	Rat "group-project2/entities/rating"
	R "group-project2/entities/room"
	U "group-project2/entities/user"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&U.Users{})
	db.AutoMigrate(&R.Rooms{})
	db.AutoMigrate(&I.Images{})
	db.AutoMigrate(&Rat.Ratings{})
}
