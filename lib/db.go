package lib

import (
	"fmt"
	"log"

	"github.com/jorgeAM/error/config"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetConnection -> get db's connection
func GetConnection() *gorm.DB {
	c := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
