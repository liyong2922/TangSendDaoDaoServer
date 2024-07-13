package base

import (
	"embed"

	"github.com/liyong2922/TangSengDaoDaoServer/modules/base/app"
	"github.com/liyong2922/TangSengDaoDaoServerLib/config"
	"github.com/liyong2922/TangSengDaoDaoServerLib/pkg/register"
)

//go:embed sql
var sqlFS embed.FS

func init() {

	register.AddModule(func(ctx interface{}) register.Module {

		return register.Module{
			SetupAPI: func() register.APIRouter {
				return app.New(ctx.(*config.Context))
			},
			SQLDir: register.NewSQLFS(sqlFS),
		}
	})
}
