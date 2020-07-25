package middleware

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//
func Log(r *ghttp.Request) {
	r.Middleware.Next()
	g.Log().Println(r.Response.Status, r.URL.Path)
}