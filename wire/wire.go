//go:build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/plh97/wire-learn/core"
	"github.com/plh97/wire-learn/router"
)

func InitWire() *gin.Engine {
	wire.Build(
		router.RouterProvider,
		core.CommonProd,
	)
	return nil
}
