package service

import (
	model "github.com/plh97/wire-learn/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewUserService(db *gorm.DB, log *zap.Logger) *UserService {
	return &UserService{db: db, log: log}
}

func (s *UserService) GetUser(id uint) (*model.User, error) {
	var user model.User
	err := s.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		s.log.Sugar().Errorf("Failed to get user with id %d: %v", id, err)
		return nil, err
	}
	return &user, nil
}
