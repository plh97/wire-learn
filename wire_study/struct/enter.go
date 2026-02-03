package _struct

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewUserService(db *gorm.DB, log *zap.Logger) *UserService {
	return &UserService{
		db:  db,
		log: log,
	}
}
type VideoService struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewVideoService(db *gorm.DB, log *zap.Logger) *VideoService {
	return &VideoService{
		db:  db,
		log: log,
	}
}

type UserApi struct {
	*UserService
	*VideoService
}
