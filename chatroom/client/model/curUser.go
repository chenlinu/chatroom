package model

import (
	"demo/chatroom/common/message"
	"net"
)

//CurUser .当前用户
type CurUser struct {
	Conn net.Conn
	message.User
}
