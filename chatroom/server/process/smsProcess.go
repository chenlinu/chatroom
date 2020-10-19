package process2

import (
	"demo/chatroom/common/message"
	"demo/chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

//SmsProcess .
type SmsProcess struct {
	//..
}

//SendGroupMsg .写方法，转发消息
func (obj *SmsProcess) SendGroupMsg(mes *message.Message) {
	//遍历 服务器端的 onlineUsers map[int]*UserProcess
	//将消息转发取出

	//取出 mes的内容 SmsMes
	var smsMsg message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal 728 err", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal 253 err", err)
		return
	}
	for ID, up := range userMgr.onlineUsers {
		if ID == smsMsg.UserID {
			continue
		}

		obj.SendMsgToEachOnlineUser(data, up.Conn)
	}
}

//SendMsgToEachOnlineUser .
func (obj *SmsProcess) SendMsgToEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个Transfer 实例 发送data

	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 168 err", err)
	}
}
