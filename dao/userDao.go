package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// 根据用户名和密码查询一条记录
func CheckUsernameAndPassword(username, password string) (*model.User, error) {
	//写sql
	sqlStr := "select id,username,password,email from users where username=? and password=?"
	//执行sql
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	scan := row.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	return user, scan
}

// 根据用户名查询一条记录
func CheckUsername(username string) (model.User, error) {
	//写sql
	sqlStr := "select id,username,password,email from users where username=?"
	//执行sql
	row := utils.Db.QueryRow(sqlStr, username)
	user := model.User{}
	scan := row.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	return user, scan
}

// 向数据库中插入用户信息
func SaveUser(user *model.User) error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	_, err := utils.Db.Exec(sqlStr, user.UserName, user.Password, user.Email)
	return err
}
