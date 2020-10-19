package process

import (
	"demo/chatroom/client/utils"
	"demo/chatroom/common/message"
	"encoding/json"
	"fmt"
)

//SmsProcess .
type SmsProcess struct {
}

//SendGroupMes 发送群聊
func (obj *SmsProcess) SendGroupMes(content string) (err error) {

	//1 创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2 创建 一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserID = curUser.UserID
	smsMes.UserStatus = curUser.UserStatus

	//3 序列化 smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes 1 json.Marshal err =", err)
		return
	}

	mes.Data = string(data)

	//4. 对mes 再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes 2 json.Marshal err =", err)
		return
	}

	//5. 将mes发送给服务器。。。
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes 2 tf.WritePkg err =", err)
		return
	}
	return
}
