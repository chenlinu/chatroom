package main

import (
	"demo/chatroom/common/message"
	process2 "demo/chatroom/server/process"
	"demo/chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

//Processor 先创建结构体
type Processor struct {
	Conn net.Conn
}

//ServerProcessMes 根据客户端发送消息种类不同，决定调用哪个函数来处理
func (obj *Processor) ServerProcessMes(mes *message.Message) (err error) {
	fmt.Println("mes.Type=", mes.Type)
	fmt.Println("mes.Type=", mes)
	switch mes.Type {
	case message.LoginMesType:
		fmt.Println("收到登录包")
		//处理登录
		up := &process2.UserProcess{
			Conn: obj.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:

		fmt.Println("收到注册包")
		//处理注册
		up := &process2.UserProcess{
			Conn: obj.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		fmt.Println("收到发送消息包")
		//处理发送消息
		//创建一个SmsProcess实例，完成转发群聊消息。
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMsg(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

//Process2 .
func (obj *Processor) Process2() (err error) {

	//循环的客户端发送的信息
	for {

		//创建一个Tansfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: obj.Conn,
		}

		//这里将读取数据包，直接封装成一个函数readpkg()
		mes, err := tf.ReadPkg()
		fmt.Println("客户端发过来的内容：", &mes)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出。。。")
				return err
			}
			fmt.Println("conn.read err =", err)
			return err
		}
		fmt.Println("由 ServerProcessMes 处理")
		err = obj.ServerProcessMes(&mes)
		if err != nil {
			return err
		}

	}
}
