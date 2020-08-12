package main

import (
	"bookstore/controller"
	"log"
	"net/http"
)

func main() {

	//设置处理静态资源 匹配
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	//修改为有session 的参数
	http.HandleFunc("/main", controller.GetPageBooksByPrice)

	// 去登录页面
	http.HandleFunc("/login", controller.Login)
	// 去注销页面
	http.HandleFunc("/logout", controller.Logout)
	// 去注册页面
	http.HandleFunc("/regist", controller.Register)

	//通过ajax的请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//获取所有图书
	//http.HandleFunc("/getBooks", controller.GetBooks)
	//获取带分页的图书信息
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	///getPageBooksByPrice
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	//添加图书
	//http.HandleFunc("/addBook", controller.AddBooks)

	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//去更新图书的页面
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	//添加或修改图书信息
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)

	serve := http.ListenAndServe(":8080", nil)
	if serve != nil {
		log.Println(serve)

	}
}
