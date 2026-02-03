package bind

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ServiceInterface interface {
	Save() string
}

type UserService struct {
	db  *gorm.DB
	log *zap.Logger
}

func (UserService) Save() string {
	return "user saved"
}

func NewUserService(db *gorm.DB, log *zap.Logger) ServiceInterface {
	return &UserService{
		db:  db,
		log: log,
	}
}

type UserApi struct {
	service *ServiceInterface
}

func NewUserApi(service ServiceInterface) *UserApi {
	return &UserApi{
		service: &service,
	}
}