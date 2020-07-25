package user_service

import (
	"errors"
	"kong1/form"
	"kong1/middleware"
	"kong1/provider/user_provider"
)

func List(data *form.UserListForm)  (_map interface{}, e error) {
	_list,_total,err := user_provider.GetListAndTotal(data)
	return map[string]interface{}{
		"data": _list,
		"total":_total,
		"user": middleware.AuthUserInfo(),
	},err
}

func Store(data *form.UserStoreForm) (e error) {

	res,err := user_provider.CheckUserExists(data.Username)
	if err != nil {
		return err
	}
	if res {
		e =  errors.New("用户已存在")
		return e
	}
	err = user_provider.Store(data)
	return err
}


func Update(data *form.UserUpdateForm)  (e error) {
	err := user_provider.Update(data)
	return err
}


func Delete(data *form.UserDeleteForm) (e error) {
	err := user_provider.Delete(data)
	return err
}

func One(data *form.UserOneForm) (_map interface{} , e error)  {
	_row, err := user_provider.One(data)
	return map[string]interface{}{
		"data": _row,
	},err
}
