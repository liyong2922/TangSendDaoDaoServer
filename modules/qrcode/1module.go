package qrcode

import (
	"github.com/liyong2922/TangSengDaoDaoServerLib/config"
	"github.com/liyong2922/TangSengDaoDaoServerLib/pkg/register"
)

func init() {

	register.AddModule(func(ctx interface{}) register.Module {

		return register.Module{
			SetupAPI: func() register.APIRouter {
				return New(ctx.(*config.Context))
			},
		}
	})
}
