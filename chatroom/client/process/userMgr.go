package process

import (
	"demo/chatroom/client/model"
	"demo/chatroom/common/message"
	"fmt"
)

//客户端维护用户map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

//当前用户，在用户登录成功后，去初始化curUser
var curUser model.CurUser

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历onlinUsers
	fmt.Println("当前在线用户列表：")
	for ID := range onlineUsers {
		fmt.Println("用户ID=\t", ID)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserID]
	if !ok { //不存在
		user = &message.User{
			UserID: notifyUserStatusMes.UserID,
		}
	}

	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserID] = user
	outputOnlineUser()
}
