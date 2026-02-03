package service

import (
	"github.com/google/wire"
	"github.com/plh97/wire-learn/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	db     *gorm.DB
	logger *zap.Logger
}

var NewUserServiceProvider = wire.NewSet(NewUserService)

func NewUserService(db *gorm.DB, logger *zap.Logger) *UserService {
	return &UserService{
		db:     db,
		logger: logger,
	}
}

func (s *UserService) GetUser(id int) ([]*model.User, error) {
	var user []*model.User
	var err error
	if id > 0 {
		err = s.db.Find(&user, id).Error
	} else {
		err = s.db.Find(&user).Error
	}
	if err != nil {
		s.logger.Error("Error fetching user", zap.Error(err))
		return nil, err
	}
	s.logger.Info("User data fetched", zap.Any("data", user))
	return user, nil
}

func (s *UserService) CreateUser(user *model.User) error {
	err := s.db.Create(user).Error
	if err != nil {
		s.logger.Error("Error creating user", zap.Error(err))
		return err
	}
	s.logger.Info("User created successfully", zap.Any("data", user))
	return nil
}
