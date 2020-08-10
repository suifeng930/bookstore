package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestSaveUser(t *testing.T) {

	fmt.Println("测试用户")
	t.Run("验证用户名或密码", testLogin)
	t.Run("验证用户名", testRegister)
	t.Run("保存用户", testSaveUser)

}

func testLogin(t *testing.T) {

	user, _ := CheckUsernameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是：", user)

}

func testRegister(t *testing.T) {
	user, _ := CheckUsername("admin")
	fmt.Println("获取用户信息是：", user)

}
func testSaveUser(t *testing.T) {

	user := &model.User{
		UserName: "admin",
		Password: "123456",
		Email:    "123",
	}
	saveUser := SaveUser(user)
	fmt.Println("创建用户成功：", saveUser)
}
