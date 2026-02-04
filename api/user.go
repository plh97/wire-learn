package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/plh97/wire-learn/model"
	"github.com/plh97/wire-learn/service"
	"github.com/plh97/wire-learn/utils"
	"go.uber.org/zap"
)

type UserApi struct {
	*zap.Logger
	*service.UserService
	*service.VideoService
}

type GetUserRequest struct {
	ID int `form:"id"`
}

var NewUserApiProvider = wire.NewSet(NewUserApi)

func NewUserApi(log *zap.Logger, userService *service.UserService, videoService *service.VideoService) *UserApi {
	return &UserApi{
		Logger:       log,
		UserService:  userService,
		VideoService: videoService,
	}
}

func (u *UserApi) GetUser(c *gin.Context) {
	var req GetUserRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := u.UserService.GetUser(req.ID)
	// u.Logger.Debug("GetUser called", zap.Any("user", req))
	u.Logger.Debug("GetUser called")
	// u.Logger.Info("GetUser called", zap.Any("user", req))
	u.Logger.Info("GetUser called")
	// u.Logger.Warn("GetUser called", zap.Any("user", req))
	u.Logger.Warn("GetUser called")
	// u.Logger.Error("GetUser called", zap.Any("request", req))
	u.Logger.Error("GetUser called")
	// u.Logger.Panic("GetUser called", zap.Any("request", req))
	// u.Logger.Fatal("GetUser called", zap.Any("request", req))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
}

func (u *UserApi) AddUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.Name = utils.RandStringRunes(6)
	user.Email = user.Name + "@google.com"
	err := u.UserService.CreateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}
