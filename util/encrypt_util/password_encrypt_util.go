package encrypt_util

import (
	"github.com/gogf/gf/crypto/gmd5"
	"math/rand"
	"time"
)

//密码加密
func PasswordEncrypt(str string,salt string)  string{
	res,_ :=  gmd5.EncryptString(str)
	res = res + salt
	res,_ = gmd5.EncryptString(res)
	return res
}

//获取加密的盐
func GeneratePasswordSalt() string {
	l := 6
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
