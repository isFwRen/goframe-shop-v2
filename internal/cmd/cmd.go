package cmd

import (
	"context"
	"goframe-shop-v2/internal/controller"
	"goframe-shop-v2/internal/controller/backend"
	"goframe-shop-v2/internal/controller/frontend"
	"goframe-shop-v2/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				return err
			}
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					//service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Admin.Create, // 管理员
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					err = gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.Rotation,     //轮播图
						controller.Position,     //手工位图
						controller.Admin.List,   // 管理员
						controller.Admin.Update, // 管理员
						controller.Admin.Delete, // 管理员
						controller.Admin.Info,   // 管理员
						controller.Data,         //数据大全
						controller.Role,         //角色
						controller.Permission,   //权限
						controller.Category,     //商品分类管理
						controller.Coupon,       //优惠券管理
						controller.Goods,        //商品管理
						controller.GoodsOptions, //商品规格管理
						//这么写是为了避免前后端重复注册相同的路由和方法
						controller.Order.List,   //订单列表
						controller.Order.Detail, //订单详情
						backend.Article,         //文章管理&CMS
					)
				})
			})
			//---------------------华丽的分割线-------------------
			// 启动前台项目gtoken
			frontendToken, err := StartFrontendGToken()
			if err != nil {
				return err
			}
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					//service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Admin.Create, // 管理员
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					err = frontendToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.User.Info,           //当前登录用户的信息
						controller.User.UpdatePassword, //当前用户修改密码
						controller.Collection,          //收藏
						controller.Praise,              //收藏
						controller.Comment,             //评论
						controller.Cart,                //购物车
						controller.Order.Add,           //下单
						controller.OrderGoodsComments,  //订单评价
						frontend.Article,               //文章 @author自愚自乐
						frontend.Refund,                //售后 @author自愚自乐
					)
				})
			})
			s.SetPort(8000)
			s.Run()
			return nil
		},
	}
)
