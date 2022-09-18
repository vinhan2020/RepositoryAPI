package database

import (
	"DemoApiPostgre/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQL struct {
	SQL *gorm.DB
}

func DBConnection(host, user, pass, port, dbName string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, pass, dbName, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_ = db.AutoMigrate(&models.User{})
	return
}
