package main

import (
	"demo/chatroom/server/model"
	"fmt"
	"net"
	"time"
)

func main() {

	//当服务器启动时，我们去初始化我们的redis的连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口监听。。。。")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	for {
		fmt.Println("等待客户端来连接服务器。。。。")

		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			continue
		}

		//一旦链接成功，则启动一个协和和客户端保持通讯
		go Process(conn)
	}
}

//Process d
func Process(conn net.Conn) {
	//
	defer conn.Close()

	//这里调用总控，创建一个
	processor := &Processor{
		Conn: conn,
	}
	err := processor.Process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协和错误 err=", err)
		return
	}
}

//初始化一个UserDao
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
