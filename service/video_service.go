package service

import (
	model "github.com/plh97/wire-learn/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VideoService struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewVideoService(db *gorm.DB, log *zap.Logger) *VideoService {
	return &VideoService{db: db, log: log}
}

func (s *VideoService) GetVideo(id uint) (*model.User, error) {
	var user model.User
	err := s.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		s.log.Sugar().Errorf("Failed to get user with id %d: %v", id, err)
		return nil, err
	}
	return &user, nil
}
