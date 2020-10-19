package process

import (
	"demo/chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMsg(msg *message.Message) {
	//显示 SmsMes
	//1. 反序列化 msg.Data
	fmt.Println("反序列化 SmsMes")
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(msg.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal 755 err", err)
		return
	}

	//显示信息
	fmt.Println("准备输出")
	info := fmt.Sprintf("用户ID：\t%d 对大家说：\t %s", smsMes.UserID, smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}
