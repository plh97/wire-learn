package user_service

import (
	"github.com/liom-source/wire-learn/models"
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

func (u *UserService) GetUser(id uint) (user *models.UserModel, err error) {
	err = u.db.Take(&user, id).Error
	if err != nil {
		if u.log != nil {
			u.log.Sugar().Warnf("GetUser error: %v", err)
		}
		return nil, err
	}
	return user, nil
}
