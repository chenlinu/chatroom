package process2

import "fmt"

//因为UserMgr实例在服务器端只有一个
//因为在很多地方都会使用到，因此我们将其定义为全局变量

var userMgr *UserMgr

//UserMgr .
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对UserMgr初始化工作

//init .
func init() {
	//1024 .
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineuser添加

//AddOnlineUser .
func (obj *UserMgr) AddOnlineUser(up *UserProcess) {
	obj.onlineUsers[up.UserID] = up
}

//DelOnlineUser .
func (obj *UserMgr) DelOnlineUser(userID int) {
	delete(obj.onlineUsers, userID)
}

//返回当前所有在线用户

//GetALLOnlineUsers .
func (obj *UserMgr) GetALLOnlineUsers() map[int]*UserProcess {
	return obj.onlineUsers
}

//根据ID返回对应的值

//GetOnlineUserByID .
func (obj *UserMgr) GetOnlineUserByID(userID int) (up *UserProcess, err error) {
	up, ok := obj.onlineUsers[userID]
	if !ok {
		//说明当前查找的用户不在线
		err = fmt.Errorf("用户%d 不存在", userID)
		return
	}
	return
}
