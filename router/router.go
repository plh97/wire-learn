package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/plh97/wire-learn/api"
)

var NewRouterProvider = wire.NewSet(NewRouter)

func NewRouter(userApi *api.UserApi, videoApi *api.VideoApi) *gin.Engine {
	r := gin.Default()
	g := r.Group("/api")
	{
		g.GET("/user", userApi.GetUser)
		g.GET("/user/add", userApi.AddUser)
	}
	{
		g.GET("/video", videoApi.GetVideo)
	}

	return r
}
