package core

import (
	"github.com/glebarez/sqlite"
	"github.com/google/wire"
	"github.com/plh97/wire-learn/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.test"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("failed to initialize logger")
	}
	return logger
}

var CommonProd = wire.NewSet(NewDB, NewLogger)
