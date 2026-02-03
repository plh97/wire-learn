package newset

import (
	"github.com/google/wire"
	"github.com/plh97/wire-learn/core"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// //////////////////// Service Layer /////////////////////
type UserService struct {
	db  *gorm.DB
	log *zap.Logger
}

var UserServiceProviderSet = wire.NewSet(NewUserService, core.NewDB, core.NewLogger)

func NewUserService(db *gorm.DB, log *zap.Logger) *UserService {
	return &UserService{
		db:  db,
		log: log,
	}
}

type GetUserRequest struct {
	ID uint `form:"id" binding:"required"`
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

type LogService struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewLogService(db *gorm.DB, log *zap.Logger) *LogService {
	return &LogService{
		db:  db,
		log: log,
	}
}

////////////////////// Service Layer /////////////////////

// /////////////////// API Layer /////////////////////
type UserApi struct {
	*UserService
	*VideoService
	*LogService
}

var UserApiProviderSet = wire.NewSet(NewUserApi, UserServiceProviderSet, NewVideoService, NewLogService)

func NewUserApi(userService *UserService, videoService *VideoService, logService *LogService) *UserApi {
	return &UserApi{
		UserService:  userService,
		VideoService: videoService,
		LogService:   logService,
	}
}

///////////////////// API Layer /////////////////////
