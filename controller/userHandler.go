package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// login  处理登录函数

func Login(w http.ResponseWriter, r *http.Request) {

	//获取前端传递过来的用户名，密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	fmt.Println("username", username)
	//调用dao 处理用户请求
	user, err := dao.CheckUsernameAndPassword(username, password)
	log.Println(user)
	if err != nil {
		log.Println("查询数据失败：")
	}
	if user.Id > 0 {
		// 用户名，密码正确
		//  1.创建一个session
		uuid := utils.CreateUUID()
		session := &model.Session{}
		session.SessionId = uuid
		session.UserId = user.Id
		session.Username = user.UserName
		//2 将session 存储到session表中
		err := dao.AddSession(session)
		if err != nil {
			log.Println("add session fails")
		}
		//  3.创建一个cookie ,让它与session关联
		cookie := http.Cookie{
			Name:     "user",
			Value:    uuid,
			HttpOnly: true,
		}
		// 4.将 cookie 发送给浏览器
		http.SetCookie(w, &cookie)
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, user)
	} else {
		//用户名 密码不正确
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		//返回前端用户信息
		t.Execute(w, "用户名或密码不正确！")
	}

}

// 用户注销处理
func Logout(w http.ResponseWriter, r *http.Request) {

	// 1.获取cookies
	cookie, err := r.Cookie("user")
	if err != nil {
		log.Println("获取cookie 失败", err.Error())
	}
	if cookie != nil {

		//获取cookie的value
		sessionId := cookie.Value
		//删除数据库中的cookie
		err := dao.DeleteSession(sessionId)
		if err != nil {
			log.Println("删除 cookie 失败 ,", err.Error())
		}
		//设置cookie失效
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)

	}

	//去首页
	GetPageBooksByPrice(w, r)

}

func Register(w http.ResponseWriter, r *http.Request) {

	//获取传递过来的参数
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	emial := r.PostFormValue("email")

	//根据用户名查询user
	user, err := dao.CheckUsername(username)
	if err != nil {
		log.Println("查询用户失败")

	}
	log.Println(user)
	if user.Id > 0 || err != nil {
		// 用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在")
	} else {
		//注册新用户
		users := &model.User{
			UserName: username,
			Password: password,
			Email:    emial,
		}
		saveUser := dao.SaveUser(users)
		if saveUser != nil {
			log.Println("插入用户数据失败了")
		}
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}

}

// 通过验证ajax 验证用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取传递过来的参数
	username := r.PostFormValue("username")

	user, err := dao.CheckUsername(username)
	if err != nil {
		log.Println("获取用户信息失败")
	}
	log.Println("获取的用户信息：", user)
	if user.Id > 0 || err != nil {
		// 用户名已存在
		str := "用户名已存在"
		w.Write([]byte(str))
	} else {
		str := "用户名可用"
		w.Write([]byte(str))
	}

}
