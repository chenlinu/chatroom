package message

//User .
type User struct {
	//要保证tag 与 json字符串的key 一致
	UserID     int    `json:"userID"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
}
