package router

import (
	"github.com/gin-gonic/gin"
	"github.com/plh97/wire-learn/api"
)

func NewUserRouter(api *api.UserApi) *gin.Engine {
	r := gin.Default()
	g := r.Group("/api")
	{
		g.GET("/user", api.GetUser)
	}
	return r
}
