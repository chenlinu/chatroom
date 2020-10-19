package message

//Message d
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //
}

//LoginMes d
type LoginMes struct {
	UserID   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	UserName string `json:"username"`
}

//LoginResMes d
type LoginResMes struct {
	Code    int    `json:"code"` //状态码，500 未注册，200登录成功
	Error   string `json:"error"`
	UsersID []int  //返回在线用户ID集合切片
}

const (
	//LoginMesType d
	LoginMesType = "LoginMes"
	//LoginResMesType s
	LoginResMesType = "LoginResMes"
	//RegisterMesType .
	RegisterMesType = "RegisterMes"
	//RegisterResMesType .
	RegisterResMesType = "RegisterResMes"
	//NotifyUserStatusMesType .
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	//SmsMesType .
	SmsMesType = "SmsMes"
)

//这里我们定义几个用户状态的常量
const (
	//
	UserOnline = iota //0
	UserOffLine
	UserBusyStatus
)

//RegisterMes .
type RegisterMes struct {
	//...
	User User `json:"user"`
}

//RegisterResMes .
type RegisterResMes struct {
	Code  int    `json:"code"` //状态码，400 已经占用，200登录成功
	Error string `json:"error"`
}

//为了配合服务器端推送用户状态变化的消息

//NotifyUserStatusMes .
type NotifyUserStatusMes struct {
	UserID int `json:"userID"`
	Status int `json:"status"`
}

//SmsMes 增加一个SMSMes //发送的消息
type SmsMes struct {
	Content string `json:"content"`
	User           //匿名结构体
}
