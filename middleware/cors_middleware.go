package middleware

import "github.com/gogf/gf/net/ghttp"

// 允许接口跨域请求
func Cors(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
