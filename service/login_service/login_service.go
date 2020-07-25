package login_service

import (
	"errors"
	"kong1/form"
	"kong1/provider/user_provider"
	"kong1/util/encrypt_util"
	"kong1/util/jwt_util"
)


// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func Login(data *form.LoginForm) (res interface{}, e error){

	row,err := user_provider.GetRowByUser(data.Username)
	if err != nil {
		return nil,err
	}
	if encrypt_util.PasswordEncrypt(data.Password,row.Salt) != row.Password {
		return nil, errors.New("用户名或密码错误")
	}

	//get jwt
	claims := jwt_util.CustomClaims{
		ID: row.Id,
		Name: row.Username,
	}
	j := jwt_util.NewJWT()
	token,er := j.CreateToken(claims)
	if er != nil {
		return nil,er
	}

	return map[string]interface{}{
		"data": row,
		"token":token,
	},err
}

//用户注册
func SignUp(data *form.SignUpForm) (res interface{}, e error) {

	if data.Password != data.PasswordConfirm {
		return nil, errors.New("两次密码不一致")
	}
	r, err := user_provider.CheckUserExists(data.Username)
	if err != nil {
		return nil, err
	}
	if r{
		return nil, errors.New("用户已存在")
	}
	row,err1 := user_provider.SignUser(data)
	if err1 != nil {
		return nil, err1
	}

	//get jwt
	claims := jwt_util.CustomClaims{
		ID: row.Id,
		Name: row.Username,
	}
	j := jwt_util.NewJWT()
	token,er := j.CreateToken(claims)
	if er != nil {
		return nil,er
	}

	return map[string]interface{}{
		"data": row,
		"token":token,
	},err
}


// 用户注销
func SignOut() error {
	return nil
}
