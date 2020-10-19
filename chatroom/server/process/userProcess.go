package process2

import (
	"demo/chatroom/common/message"
	"demo/chatroom/server/model"
	"demo/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

//UserProcess .
type UserProcess struct {
	//字段？
	Conn net.Conn

	//增加一个字段，表示该Conn是哪个用户
	UserID int
}

//NotifyOthersOnlineUser .这里编写通知所有在线用户的方法 ,传入自己的ID
func (obj *UserProcess) NotifyOthersOnlineUser(userID int) {
	//遍历 onlineusers,然后一个一个的发送 notifyuserstatusmes
	//userMap := userMgr.GetALLOnlineUsers()
	for id, up := range userMgr.onlineUsers {
		if id == userID {
			continue //过滤自己
		}

		//开始通知
		up.NotifyMeOnline(userID)
	}

}

//NotifyMeOnline .
func (obj *UserProcess) NotifyMeOnline(userID int) {
	//组装我们的消息 NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserID = userID
	notifyUserStatusMes.Status = message.UserOnline

	//将 notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal 1 err = ", err)
		return
	}

	//将序列化后的notifyUserStatusMes赋值给 mes.Data
	mes.Data = string(data)

	//对mes再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal 2 err = ", err)
		return
	}

	//发送,创建我们Transfer实例，发送
	tf := &utils.Transfer{
		Conn: obj.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifyMeOnline err = ", err)
		return
	}

}

//ServerProcessLogin 编写一个函数serverProcessLogin函数
func (obj *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2声明一个LoginResMes，并完成赋值
	var loginResMes message.LoginResMes

	//去redis完成验证
	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPwd)
	if err != nil {
		if err == model.ErrorUserNotExists {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ErrorUserPassword {
			loginResMes.Code = 300
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200

		//这里，因为用户登录成功，要把信息放到OnlineUsers里
		//另外将用户的UserID 赋给 obj.UserID
		obj.UserID = loginMes.UserID
		userMgr.AddOnlineUser(obj)
		obj.NotifyOthersOnlineUser(loginMes.UserID)

		//将当前在线用户的ID 放到LoginResMes.UsersID
		//遍历 userMgr.onlinUsers
		for ID := range userMgr.onlineUsers {
			loginResMes.UsersID = append(loginResMes.UsersID, ID)
		}
		fmt.Println(user, "登录成功！")
	}

	//3 将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	//4.将 data 赋值给 resMes
	resMes.Data = string(data)

	//5 对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail err=", err)
		return
	}

	//6 ,发送data MVC
	tf := &utils.Transfer{
		Conn: obj.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//ServerProcessRegister .
func (obj *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	var resMsg message.Message
	resMsg.Type = message.RegisterResMesType
	var RegisterResMes message.RegisterResMes

	fmt.Println("尝试注册用户=", &registerMes.User)
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ErrorUserExists {
			fmt.Println("code = 505")
			RegisterResMes.Code = 505
			RegisterResMes.Error = model.ErrorUserExists.Error()
		} else {
			fmt.Println("code = 506")
			RegisterResMes.Code = 506
			RegisterResMes.Error = "注册发生未知错误"
		}
	} else {
		fmt.Println("code = 200")
		RegisterResMes.Code = 200
	}

	fmt.Println("序列化注册结果")
	data, err := json.Marshal(RegisterResMes)
	if err != nil {
		fmt.Println("json Marshal err =", err)
		return
	}

	//5
	resMsg.Data = string(data)

	//6
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	//这里需要处理服务器返回的信息
	//创建一个Tansfer实例
	tf := &utils.Transfer{
		Conn: obj.Conn,
	}
	fmt.Println("向客户端返回注册结果")
	//发送data
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册信息发送失败 err =", err)
	}

	return
}
