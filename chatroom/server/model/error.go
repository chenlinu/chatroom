package model

import "errors"

//根据业务逻辑需要，自定义一些错误

var (
	//ErrorUserNotExists .
	ErrorUserNotExists = errors.New("用户不存在。。")
	//ErrorUserExists .
	ErrorUserExists = errors.New("用户已经存在。。")
	//ErrorUserPassword .
	ErrorUserPassword = errors.New("密码不正确 ")
)
