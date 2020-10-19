package utils

import (
	"demo/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//Transfer 将方法关联到结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //这是传输时，使用缓冲
}

//WritePkg .
func (obj *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte

	binary.BigEndian.PutUint32(obj.Buf[:4], pkgLen)
	n, err := obj.Conn.Write(obj.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//发送data本身
	n, err = obj.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}

//ReadPkg .
func (obj *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据。。。server.exe")

	//conn.read 只有在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭了conn，就不会阻塞
	_, err = obj.Conn.Read(obj.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(obj.Buf[0:4])

	fmt.Println("客户端发送 pkgLen =", pkgLen)

	//从conn里读取pkgLen个字节到buf里
	n, err := obj.Conn.Read(obj.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return
	}

	//反序列化
	err = json.Unmarshal(obj.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}

	return
}
