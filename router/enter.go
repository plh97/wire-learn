package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liom-source/wire-learn/api/user_api"
)

func NewRouter(userApi *user_api.UserApi) *gin.Engine {
	r := gin.Default()
	g := r.Group("/api")
	{
		g.GET("/user", userApi.GetUserView)
	}
	return r
}
