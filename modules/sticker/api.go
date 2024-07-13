package sticker

import (
	"errors"

	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/config"
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/pkg/log"
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/pkg/wkhttp"
)

type Sticker struct {
	ctx *config.Context
	log.Log
}

// Route 配置路由规则
func (r *Sticker) Route(l *wkhttp.WKHttp) {
	v := l.Group("/v1/sticker")
	{
		v.GET("/user", r.user)
	}
}

func (rb *Sticker) user(c *wkhttp.Context) {
	c.ResponseError(errors.New("数据格式有误！"))
	//111
}
