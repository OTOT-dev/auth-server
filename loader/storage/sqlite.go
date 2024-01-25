package storage

import (
	"auth-server/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dbFileName = "auth-server.db"
	db         *gorm.DB
)

func init() {
	err := initStorageEngine()
	if err != nil {
		log.Fatalf("init storage Engine error: %s", err)
	}
}

type EngineStorage struct{}

func (EngineStorage) GetStorageDB() *gorm.DB {
	return db
}

func initStorageEngine() (err error) {
	db, err = gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(model.UserProps{})
	if err != nil {
		return
	}
	return
}
