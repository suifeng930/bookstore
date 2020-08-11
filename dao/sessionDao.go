package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"log"
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
