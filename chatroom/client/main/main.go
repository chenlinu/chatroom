package main

import (
	"demo/chatroom/client/process"
	"fmt"
)

var userID int
var userPwd string
var nickName string

func main() {
	var key int
	//判断是否继续显示菜单

	for true {
		fmt.Println("================欢迎登录多人聊天系统~=================")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册账号")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3)")

		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入UserId")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入UserPwd")
			fmt.Scanf("%s\n", &userPwd)
			//完成登录
			//1创建一个userprocess的实例
			up := &process.UserProcess{}
			up.Login(userID, userPwd)
		case 2:
			fmt.Println("注册账号")
			fmt.Println("请输入UserId")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入UserPwd")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入NickName")
			fmt.Scanf("%s\n", &nickName)
			up := &process.UserProcess{}
			err := up.Register(userID, userPwd, nickName)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("不正确，请重新输入")
		}
	}

}
