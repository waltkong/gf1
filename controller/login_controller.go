package controller

import (
	"github.com/gogf/gf/net/ghttp"
	"kong1/form"
	"kong1/service/login_service"
	"kong1/util/error_util"
	"kong1/util/response_util"
)

type LoginController struct {

}

// 登录接口
func (c *LoginController) Login(r *ghttp.Request) {
	// 输入参数检查
	data,err := form.LoginFormCheck(r)
	error_util.CommonCheckError(r,err)

	res,err2 := login_service.Login(data)
	error_util.CommonCheckError(r,err2)

	response_util.JsonExit(r, 0, "ok",res)
}

// 注册接口
func (c *LoginController) SignUp(r *ghttp.Request) {
	// 输入参数检查
	data,err := form.SignUpFormCheck(r)
	error_util.CommonCheckError(r,err)
	//逻辑层
	res,err2 := login_service.SignUp(data)
	error_util.CommonCheckError(r,err2)
	//响应
	response_util.JsonExit(r, 0, "ok",res)
}

//注销接口
func (c *LoginController) SignOut(r *ghttp.Request)  {

	err := login_service.SignOut()
	error_util.CommonCheckError(r,err)

	response_util.JsonExit(r, 0, "ok")
}
