package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"todos-rest/model/entity"
)

func NewDb() *gorm.DB {
	dbUrl := viper.GetString("DB_URL")
	db, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		panic("failed to connect database")
	}

	db.CreateTable(&entity.Todo{TodoFields: &entity.TodoFields{}})
	db.Debug().AutoMigrate(&entity.Todo{})
	return db
}
