package model

import (
	"demo/chatroom/common/message"
	"encoding/json"
	"fmt"
	"redigo-master/redis"
)

//我们在服务器启动时，就初始化一个userDao实例
//把它做成全局的变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

//定义一个UserDao结构体

//UserDao .
type UserDao struct {
	pool *redis.Pool
}

//NewUserDao .
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (obj *UserDao) getUserByID(conn redis.Conn, ID int) (user *message.User, err error) {

	//通过ID 去 redis 查询这个用户
	res, err := redis.String(conn.Do("Hget", "users", ID))
	if err != nil {
		if err == redis.ErrNil { //表示在users哈希中，没有找到对应ID
			err = ErrorUserNotExists
		}
		return
	}

	user = &message.User{}

	// 这里需要把res 反序列化 成User实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)

		return
	}
	return
}

//完成登录的校验

//Login .完成对用户的验证
func (obj *UserDao) Login(userID int, userPwd string) (user *message.User, err error) {
	//从UserDao的连接池中取出一根连接
	conn := obj.pool.Get()
	defer conn.Close()
	user, err = obj.getUserByID(conn, userID)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ErrorUserPassword
	}
	return
}

//Register .完成对用户的验证
func (obj *UserDao) Register(user *message.User) (err error) {
	//从UserDao的连接池中取出一根连接
	conn := obj.pool.Get()
	defer conn.Close()
	_, err = obj.getUserByID(conn, user.UserID)
	if err == nil {
		fmt.Println("已存在")
		err = ErrorUserExists
		return
	}
	//这时，说明UserID还没有被注册，可以注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	//入库
	_, err = conn.Do("HSet", "users", user.UserID, string(data))
	if err != nil {
		fmt.Println("注册用户错误 err = ", err)
		return
	}
	return
}
