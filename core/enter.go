package core

import (
	"github.com/glebarez/sqlite"
	"github.com/google/wire"
	"github.com/plh97/wire-learn/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var CommonProvider = wire.NewSet(NewDB, NewLogger)

func NewDB() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&model.User{}, &model.Video{})

	return DB
}

func NewLogger() *zap.Logger {
	Logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return Logger
}
