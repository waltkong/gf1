package controller

import (
	"github.com/gogf/gf/net/ghttp"
	"kong1/form"
	"kong1/service/user_service"
	"kong1/util/error_util"
	"kong1/util/response_util"
)

type UserController struct {

}


// 用户列表
func (c *UserController) List(r *ghttp.Request) {
	// 输入参数检查
	data,err := form.UserListFormCheck(r)
	error_util.CommonCheckError(r,err)

	//逻辑层
	res,err2 := user_service.List(data)
	error_util.CommonCheckError(r,err2)

	response_util.JsonExit(r, 0, "ok",res)
}

//用户新增
func (c *UserController) Store(r *ghttp.Request)  {
	// 输入参数检查
	data,err := form.UserStoreFormCheck(r)
	error_util.CommonCheckError(r,err)

	//逻辑层
	err2 := user_service.Store(data)
	error_util.CommonCheckError(r,err2)

	response_util.JsonExit(r, 0, "ok")
}

//用户更新
func (c *UserController) Update(r *ghttp.Request)  {
	// 输入参数检查
	data,err := form.UserUpdateFormCheck(r)
	error_util.CommonCheckError(r,err)

	//逻辑层
	err2 := user_service.Update(data)
	error_util.CommonCheckError(r,err2)

	response_util.JsonExit(r, 0, "ok")
}

//用户删除
func (c *UserController) Delete(r *ghttp.Request)  {
	// 输入参数检查
	data,err := form.UserDeleteFormCheck(r)
	error_util.CommonCheckError(r,err)

	//逻辑层
	err2 := user_service.Delete(data)
	error_util.CommonCheckError(r,err2)

	response_util.JsonExit(r, 0, "ok")
}

//获取指定id的用户
func (c *UserController) One(r *ghttp.Request)  {
	// 输入参数检查
	data,err := form.UserOneFormCheck(r)
	error_util.CommonCheckError(r,err)

	//逻辑层
	res,err2 := user_service.One(data)
	error_util.CommonCheckError(r,err2)

	response_util.JsonExit(r, 0, "ok",res)
}