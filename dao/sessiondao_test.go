package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestSaveSession(t *testing.T) {

	fmt.Println("测试用户")
	t.Run("测试添加session", testAddSession)
	//t.Run("删除session", TestDeleteSession)

}

func TestDeleteSession(t *testing.T) {
	user := DeleteSession("132123")
	fmt.Println("删除session成功：", user)

}
func testAddSession(t *testing.T) {

	user := &model.Session{
		Username:  "admin",
		SessionId: "132123",
		UserId:    1,
	}
	saveUser := AddSession(user)
	fmt.Println("创建session成功：", saveUser)
}
