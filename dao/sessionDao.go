package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"log"
	"net/http"
)

// add Session   向数据库中添加session
func AddSession(session *model.Session) error {

	//sql
	sqlStr := "insert into sessions values (?,?,?)"

	//执行
	_, err := utils.Db.Exec(sqlStr, session.SessionId, session.Username, session.UserId)
	if err != nil {
		log.Println(" add session fails :", err.Error())
		return err
	}
	return nil
}

//  delete session
func DeleteSession(sessionId string) error {

	sqlStr := "delete from sessions where session_id =?"
	_, err := utils.Db.Exec(sqlStr, sessionId)
	if err != nil {
		log.Println("delete error fails ", err.Error())
		return err
	}
	return nil

}

// get Session
func GetSessionById(sessionId string) (*model.Session, error) {

	sqlStr := "select session_id,username,user_id from sessions where session_id=?"

	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		log.Println("select session is fail", err.Error())
		return nil, err
	}
	row := stmt.QueryRow(sessionId)
	session := &model.Session{}
	row.Scan(&session.SessionId, &session.Username, &session.UserId)
	return session, nil

}

// 判断用户登录的方法 ，加入cookie
// false  没有登录
// true   用户已经登录
func IsLogin(r *http.Request) (bool, string) {
	// 1. 获取前台传递过来的cookie
	cookies, err := r.Cookie("user")
	if err != nil {
		log.Println("get cookie is fail ", err.Error())
	}
	log.Println("cookies:", cookies)
	if cookies != nil {
		// 获取cookid values
		sessionId := cookies.Value
		//去数据库中根据sessionId 查询session
		session, err := GetSessionById(sessionId)
		if err != nil {
			log.Println("获取session 失败，", err.Error())
		}
		log.Println("session :", session)
		if session.UserId > 0 {
			//已经登录
			return true, session.Username
		}
	}

	return false, ""

}
