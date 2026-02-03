package build

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Redis struct {
}

func NewRedis() *Redis {
	return &Redis{}
}

// //////////////////// Service Layer /////////////////////
type UserService struct {
	db    *gorm.DB
	log   *zap.Logger
	redis *Redis
}

func NewUserService(db *gorm.DB, log *zap.Logger, redis *Redis) *UserService {
	return &UserService{
		db:    db,
		log:   log,
		redis: redis,
	}
}

type GetUserRequest struct {
	ID uint `form:"id" binding:"required"`
}

////////////////////// Service Layer /////////////////////

// /////////////////// API Layer /////////////////////
type UserApi struct {
	*UserService
}

func NewUserApi(userService *UserService) *UserApi {
	return &UserApi{
		UserService: userService,
	}
}

///////////////////// API Layer /////////////////////

type Info struct {
}

func NewInfo(db *gorm.DB) (*Info, error) {
	return &Info{}, nil
}
