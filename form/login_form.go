package form

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type LoginForm struct {
	Username  string `p:"username"  v:"required|length:3,30#请输入账号|账号长度为:min到:max位"`
	Password  string `p:"password"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
}


type SignUpForm struct {
	Username  string `p:"username"  v:"required|length:3,30#请输入账号|账号长度为:min到:max位"`
	Password  string `p:"password"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	PasswordConfirm string `p:"password_confirm"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
}


func LoginFormCheck(r *ghttp.Request)  (*LoginForm, error) {
	var data *LoginForm
	// 输入参数检查
	if err := r.GetStruct(&data); err != nil {
		return nil, err
	}
	if err := gvalid.CheckStruct(data, nil); err != nil {
		return nil, err
	}
	return data,nil
}

func SignUpFormCheck(r *ghttp.Request)  (*SignUpForm, error) {
	var data *SignUpForm
	// 输入参数检查
	if err := r.GetStruct(&data); err != nil {
		return nil, err
	}
	if err := gvalid.CheckStruct(data, nil); err != nil {
		return nil, err
	}
	return data,nil
}