package user_provider

import (
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"kong1/form"
	"kong1/model"
	"kong1/util/encrypt_util"
	"time"
)

//获取用户列表和总数
func GetListAndTotal(data *form.UserListForm) (r gdb.Result,total int ,err error) {
	o := new(model.User).GetInstance()
	if data.Starttime != ""{
		o = o.Where("createtime>",data.Starttime)
	}
	if data.Endtime != ""{
		o = o.Where("createtime<",data.Endtime)
	}
	if data.Username != ""{
		o = o.Where("username like ?", "%"+data.Username+"%")
	}

	_commonListData := form.GetCommonListFormData(&data.CommonListFormData)

	offset := (_commonListData.PageIndex - 1) * _commonListData.EachPage
	order := _commonListData.SortBy + " " + _commonListData.SortWay

	r,err = o.Limit(offset,_commonListData.EachPage).Order(order).FindAll()
	if err != nil {
		return nil, 0, err
	}

	total,err = o.Count()

	return
}

//保存用户数据
func Store(data *form.UserStoreForm)  (e error){
	o := new(model.User).GetInstance()
	salt := encrypt_util.GeneratePasswordSalt()
	password := encrypt_util.PasswordEncrypt(data.Password,salt)
	_map := g.Map{
		"username": data.Username,
		"password": password,
		"salt": salt,
		"phone": data.Phone,
		"email": data.Email,
		"nickname": data.Nickname,
		"createtime":time.Now().Unix(),
	}
	_, err := o.Data(_map).Insert()
	return err
}

// 更新用户数据
func Update(data *form.UserUpdateForm) (e error){
	if data.Id < 1 {
		e =  errors.New("id 非法")
		return e
	}
	o := new(model.User).GetInstance()
	_map := g.Map{
		"nickname": data.Nickname,
		"phone": data.Phone,
		"email": data.Email,
	}
	_, e = o.Data(_map).Where("id", data.Id).Update()
	return e
}

//删除
func Delete(data *form.UserDeleteForm) (e error){
	o := new(model.User).GetInstance()
	_, err := o.Where("id", data.Id).Delete()
	return err
}

//获取一个
func One(data *form.UserOneForm) (res gdb.Record, e error)  {
	o := new(model.User).GetInstance()
	r,err := o.Where("id", data.Id).FindOne()
	return r,err
}

//检测用户存在
func CheckUserExists(u string)  (bool ,error) {
	o := new(model.User).GetInstance()
	r,err := o.Where("username", u).FindOne()
	if r != nil {
		return true,err
	}
	return false,err
}

//通过username获取row
func GetRowByUser(u string)  (r *model.User, e error) {
	var user *model.User
	o := new(model.User).GetInstance()
	err := o.Where("username", u).Scan(&user)
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	return user,err
}

//注册
func SignUser(data *form.SignUpForm) (res *model.User, e error) {
	o := new(model.User).GetInstance()
	salt := encrypt_util.GeneratePasswordSalt()
	password := encrypt_util.PasswordEncrypt(data.Password,salt)
	user := &model.User{
		Username: data.Username,
		Password: password,
		Salt: salt,
		Createtime: (int) (time.Now().Unix()),
	}
	_, err := o.Data(user).Insert()
	return user,err
}