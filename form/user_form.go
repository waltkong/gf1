package form

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type UserListForm struct {
	Username  string `p:"username"`
	Starttime string `p:"starttime"`
	Endtime string `p:"endtime"`

	CommonListFormData CommonListForm
}

type UserStoreForm struct {
	Username  string `p:"username"  v:"required|length:3,30#请输入账号|账号长度为:min到:max位"`
	Password  string `p:"password"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Nickname  string `p:"nickname"`
	Email  string `p:"email"`
	Phone  string `p:"phone"`
}

type UserUpdateForm struct {
	Id int `p:"id"  v:"required"`
	Nickname  string `p:"nickname"`
	Email  string `p:"email"`
	Phone  string `p:"phone"`
}

type UserOneForm struct {
	Id int `p:"id"  v:"required"`
}

type UserDeleteForm struct {
	Id int `p:"id"  v:"required"`
}

// 用户列表参数校验
func UserListFormCheck(r *ghttp.Request)  (*UserListForm , error){
	var data *UserListForm
	// 输入参数检查
	if err := r.GetStruct(&data); err != nil {
		return nil, err
	}
	if err := gvalid.CheckStruct(data, nil); err != nil {
		return nil, err
	}
	return data,nil
}

func UserStoreFormCheck(r *ghttp.Request) (*UserStoreForm, error) {
	var data *UserStoreForm
	// 输入参数检查
	if err := r.GetStruct(&data); err != nil {
		return nil, err
	}
	if err := gvalid.CheckStruct(data, nil); err != nil {
		return nil, err
	}
	return data,nil
}

func UserUpdateFormCheck(r *ghttp.Request)  (*UserUpdateForm, error) {
	var data *UserUpdateForm
	// 输入参数检查
	if err := r.GetStruct(&data); err != nil {
		return nil, err
	}
	if err := gvalid.CheckStruct(data, nil); err != nil {
		return nil, err
	}
	return data,nil
}

func UserOneFormCheck(r *ghttp.Request)  (*UserOneForm, error) {
	var data *UserOneForm
	// 输入参数检查
	if err := r.GetStruct(&data); err != nil {
		return nil, err
	}
	if err := gvalid.CheckStruct(data, nil); err != nil {
		return nil, err
	}
	return data,nil
}

func UserDeleteFormCheck(r *ghttp.Request)  (*UserDeleteForm, error) {
	var data *UserDeleteForm
	// 输入参数检查
	if err := r.GetStruct(&data); err != nil {
		return nil, err
	}
	if err := gvalid.CheckStruct(data, nil); err != nil {
		return nil, err
	}
	return data,nil
}
