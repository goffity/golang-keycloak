package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitMySQLDatabase() (*gorm.DB, error) {
	dsn := "host=database user=user1 password=p@ssW0rd dbname=identity port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
