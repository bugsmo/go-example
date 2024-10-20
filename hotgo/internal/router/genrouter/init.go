package genrouter

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/service"
	"hotgo/utility/simple"

	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	NoLoginRouter       []interface{} // 无需登录
	LoginRequiredRouter []interface{} // 需要登录
)

// Register 注册通过代码生成的后台路由
func Register(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(simple.RouterPrefix(ctx, consts.AppAdmin), func(group *ghttp.RouterGroup) {
		if len(NoLoginRouter) > 0 {
			group.Bind(NoLoginRouter...)
		}
		group.Middleware(service.Middleware().AdminAuth)

		if len(LoginRequiredRouter) > 0 {
			group.Bind(LoginRequiredRouter...)
		}
	})
}
