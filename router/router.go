package router

import (

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"kong1/controller"
	"kong1/middleware"
)

func init()  {
	s := g.Server()
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")

	s.BindHandler("/hello", func(r *ghttp.Request){
		r.Response.Writeln("hello")
	})


	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Cors)

		_loginController := new(controller.LoginController)
		group.ALL("/login", _loginController)

		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Auth)

			_userController := new(controller.UserController)
			group.ALL("/user", _userController)

		})
	})
}
