package openapi

import (
	_ "embed"

	"github.com/liyong2922/TangSengDaoDaoServerLib/config"
	"github.com/liyong2922/TangSengDaoDaoServerLib/pkg/register"
)

//go:embed swagger/api.yaml
var swaggerContent string

func init() {
	register.AddModule(func(ctx interface{}) register.Module {
		x := ctx.(*config.Context)
		api := New(x)
		return register.Module{
			Name:    "openapi",
			Swagger: swaggerContent,
			SetupAPI: func() register.APIRouter {
				return api
			},
		}
	})
}
