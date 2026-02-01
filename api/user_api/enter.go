package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/liom-source/wire-learn/service/user_service"
)

type UserApi struct {
	userService *user_service.UserService
}

type GetUserRequest struct {
	ID uint `form:"id" binding:"required"`
}

func NewUserApi(userService *user_service.UserService) *UserApi {
	return &UserApi{
		userService: userService,
	}
}

func (h *UserApi) GetUserView(c *gin.Context) {
	var req GetUserRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(200, gin.H{"code": -1, "msg": "Invalid request"})
		return
	}

	user, err := h.userService.GetUser(req.ID)
	if err != nil {
		c.JSON(200, gin.H{"code": -1, "msg": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": 0, "msg": "Success", "data": user})
}
