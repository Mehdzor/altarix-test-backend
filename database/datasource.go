package database

import (
	"github.com/jinzhu/gorm"
	"altarix_test/model"
	log "github.com/Sirupsen/logrus"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DataSource *gorm.DB

func InitDB() {
	var err error
	DataSource, err = gorm.Open("postgres", "host=localhost dbname=altarix sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	migrateDB()
}

func migrateDB() {
	DataSource.AutoMigrate(&model.Product{})
}