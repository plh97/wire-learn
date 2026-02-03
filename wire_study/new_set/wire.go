//go:build wireinject

package newset

import (
	"github.com/google/wire"
	"github.com/plh97/wire-learn/core"
)

func InitUserApi() *UserApi {
	wire.Build(
		UserApiProviderSet,
		// core.NewDB,
		// core.NewLogger,
	)
	return nil
}
