package database

import (
	"Eight/src/internal/configs"
	"Eight/src/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewClient(dbModel configs.DbInitModel) *gorm.DB {
	dsn := "host=" + dbModel.DbHost +
		" user=" + dbModel.DbUser +
		" password=" + dbModel.DbPassword +
		" dbname=" + dbModel.DbName +
		" port=" + dbModel.DbPort +
		" sslmode=disable"

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	dbMigrateErr := db.AutoMigrate(&domain.User{})
	if dbMigrateErr != nil {
		log.Fatal(dbMigrateErr)
	}

	return db
}
