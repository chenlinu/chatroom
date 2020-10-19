package process

import (
	"demo/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"demo/chatroom/client/utils"
)

//UserProcess .
type UserProcess struct {
	//字段...
	//Conn net.Conn
}

//Register .
func (obj *UserProcess) Register(userID int,
	userPwd string, nickName string) (err error) {
	// 1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("client conn err", err)
		return
	}

	//延时关闭
	defer conn.Close()
	var msg message.Message
	msg.Type = message.RegisterMesType

	//3 创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserID = userID
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = nickName

	//4 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	fmt.Println("准备注册包")
	//5
	msg.Data = string(data)

	//6
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	//这里需要处理服务器返回的信息
	//创建一个Tansfer实例

	fmt.Println("创建 Transfer")
	tf := &utils.Transfer{
		Conn: conn,
	}

	//发送data给服务器端
	fmt.Println("准备注册包a")
	err = tf.WritePkg(data)
	fmt.Println("准备注册包b")
	if err != nil {
		fmt.Println("注册信息发送失败 err =", err)
	}

	fmt.Println("读返回包a")
	mes, err := tf.ReadPkg()
	fmt.Println("读返回包b")
	if err != nil {
		fmt.Println("readPkg(conn) fail", err)
		return
	}

	fmt.Println("反序列化服务器返回的包")
	//将mes的data部分反序列化
	var regitserResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &regitserResMes)
	if regitserResMes.Code == 200 {
		fmt.Println("注册成功，可以登录一下试试")
		os.Exit(0)
	} else if regitserResMes.Code == 500 {
		fmt.Println(regitserResMes.Error)
		os.Exit(0)
	} else {
		fmt.Printf("注册失败。Code = %v, res =%v ", regitserResMes.Code, regitserResMes.Error)
		os.Exit(0)
	}
	return
}

//Login .
func (obj *UserProcess) Login(userID int, userPwd string) (err error) {
	/* 	fmt.Printf("登录信息： %v,%v ", userID, userPwd)
	   	return nil */

	// 1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("client conn err", err)
		return
	}

	//延时关闭
	defer conn.Close()

	var msg message.Message
	msg.Type = message.LoginMesType

	//3 创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json. err=", err)
		return
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json err =", err)
		return
	}

	//7到这个时候，data就是我们要发送的消息
	//先把data长度发送给服务器
	//把长度转成切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte

	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	fmt.Printf("客户端发送=%s", string(data))

	_, err = conn.Write(data)
	if n != 4 || err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	//这里需要处理服务器返回的信息
	//创建一个Tansfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err := tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg(conn) fail", err)
		return
	}
	//将mes的data部分反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化curUser
		curUser.Conn = conn
		curUser.UserID = userID
		curUser.UserStatus = message.UserOnline //状态

		//fmt.Println("登录成功")

		//可以显示当前在线用户列表
		//遍历 loginResMes.UsersID
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersID {
			if v == userID {
				continue
			}
			fmt.Println("用户ID：\t", v)
			//完成 客户端的 onlineUsers 完成初始化
			user := &message.User{
				UserID:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println("")

		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器通讯
		go serverProcessMes(conn)
		//1 显示登录成功的菜单
		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
