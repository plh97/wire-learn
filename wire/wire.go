//go:build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/plh97/wire-learn/api"
	"github.com/plh97/wire-learn/core"
	"github.com/plh97/wire-learn/router"
	"github.com/plh97/wire-learn/service"
)

func InitWire() *gin.Engine {
	wire.Build(
		core.CommonProvider,
		api.NewUserApiProvider,
		api.NewVideoApiProvider,
		router.NewRouterProvider,
		service.NewUserServiceProvider,
		service.NewVideoServiceProvider,
	)
	return nil
}
