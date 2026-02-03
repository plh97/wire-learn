package bind

import (
	"github.com/google/wire"
	"github.com/plh97/wire-learn/core"
)

func InitUserService() ServiceInterface {
	wire.Build(
		NewUserService,
		core.NewDB,
		core.NewLogger,
	)
	return nil
}
