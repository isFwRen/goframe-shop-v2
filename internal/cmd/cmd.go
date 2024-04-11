package cmd

import (
	"context"
	"goframe-shop-v2/internal/controller"
	"goframe-shop-v2/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"goframe-shop-v2/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					//service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					hello.NewV1(),
					controller.Rotation, //轮播图
					controller.Position, //手工位图
					controller.Admin,    //手工位图
					controller.Login,    //登录
				)
			})
			s.Run()
			return nil
		},
	}
)
