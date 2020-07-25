package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"kong1/model"
	"kong1/util/jwt_util"
	"kong1/util/response_util"
	"strings"
)

const AUTH_NAME = "auth_user_info"

var  authUserInfo  *model.User

// 鉴权中间件，只有登录成功之后才能通过
func Auth(r *ghttp.Request) {
	token,ok := r.Get("token").(string)
	if !ok || token == ""{
		authStr := r.Header.Get("Authorization")
		if s := strings.Split(authStr, " "); len(s) == 2 {
			token = s[1]
		}
	}
	j := jwt_util.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		response_util.JsonExit(r,403,err.Error())
	}

	r.SetCtxVar(AUTH_NAME, claims)
	setAuthUserInfo(r,claims)

	r.Middleware.Next()
}

func setAuthUserInfo(r *ghttp.Request,claims *jwt_util.CustomClaims) {
	var user *model.User
	o := new(model.User).GetInstance()

	err := o.Where("id", claims.ID).Scan(&user)
	if user == nil || err != nil {
		response_util.JsonExit(r,403,"auth error")
	}
	authUserInfo = user
}

//获取认证用户信息
func AuthUserInfo()  *model.User {
	return authUserInfo
}