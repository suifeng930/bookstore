package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// indexHandler  去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//获取页码信息
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//获取分页的books
	page, err := dao.GetPageBooks(pageNo)

	if err != nil {
		log.Println("获取图书失败")
	}
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)

}

//get books 获取所有图书

func GetPageBooks(w http.ResponseWriter, r *http.Request) {

	//获取页码信息
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//获取分页的books
	page, err := dao.GetPageBooks(pageNo)

	if err != nil {
		log.Println("获取图书失败111")
	}
	//解析模板文件views/pages/manager/book_manager.html
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))

	//执行
	t.Execute(w, page)

}

//get books 获取所有图书

func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {

	//获取页码信息
	pageNo := r.FormValue("pageNo")

	// 获取form表单传递过来的数据
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	var page *model.Page
	var err error
	if pageNo == "" {
		pageNo = "1"
	}
	if minPrice != "" && maxPrice != "" {

		page, err = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		//将价格参数信息传递到前台页面
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	} else {
		log.Println("前台传递的参数信息有误")
		//获取分页的books
		page, err = dao.GetPageBooks(pageNo)
	}

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
		session, err := dao.GetSessionById(sessionId)
		if err != nil {
			log.Println("获取session 失败，", err.Error())
		}
		log.Println("session :", session)
		if session.UserId != 0 {
			//获取到session -----》 已经登录了
			page.IsLogin = true
			page.UserName = session.Username

		}

	}
	flag, username := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.UserName = username.Username
	}

	log.Println("page:", page)
	if err != nil {
		log.Println("获取图书失败222")
	}
	//解析模板文件views/pages/manager/book_manager.html
	t := template.Must(template.ParseFiles("views/index.html"))

	//执行
	t.Execute(w, page)

}

//get books 获取所有图书
//
//func GetBooks(w http.ResponseWriter, r *http.Request) {

//	//调用bookdao 获取所有图书的信息
//	books, err := dao.GetBooks()
//	if err != nil {
//		log.Println("获取图书失败")
//	}
//	//解析模板文件views/pages/manager/book_manager.html
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//
//	//执行
//	t.Execute(w, books)
//
//}

// 添加一本图书
func AddBooks(w http.ResponseWriter, r *http.Request) {
	//获取前台页面传递的参数信息
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	//将价格、销量课库存进行转换
	fPrice, err := strconv.ParseFloat(price, 64)
	iSales, err := strconv.ParseInt(sales, 10, 0)
	iStock, err := strconv.ParseInt(stock, 10, 0)
	if err != nil {
		log.Println(" dataFormat is fail ", err)
	}
	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	//调用dao addBooks
	err = dao.AddBook(book)
	if err != nil {
		log.Println("add book is fail :", err)
	}
	//调用getbooks处理器函数 再查询一次数据库，
	GetPageBooks(w, r)
}

//删除指定图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//获取前台传递过来的数据
	bookId := r.FormValue("bookId")
	id, _ := strconv.Atoi(bookId)
	err := dao.DeleteBook(id)
	if err != nil {
		log.Println("delete book is fail :", err.Error())
	}
	//返回到bookManager.html
	GetPageBooks(w, r)
}

//去更新图书信息
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	//获取更新图书的id
	bookId := r.FormValue("bookId")
	if bookId != "" {
		id, _ := strconv.Atoi(bookId)
		//调用bookdao中获取图书的函数
		book, err := dao.GetBookById(id)
		if err != nil {
			log.Println("get books is fail :", err.Error())

		}
		//去修改图书的页面
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, book)

	} else {
		//去添加图书页面
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, "")
	}
}

//更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	//获取前台页面传递的参数信息
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	bookId := r.PostFormValue("bookId")
	//将价格、销量课库存进行转换
	fPrice, err := strconv.ParseFloat(price, 64)
	iSales, err := strconv.ParseInt(sales, 10, 0)
	iStock, err := strconv.ParseInt(stock, 10, 0)
	id, err := strconv.ParseInt(bookId, 10, 0)
	if err != nil {
		log.Println(" dataFormat is fail ", err)
	}
	book := &model.Book{
		Id:      int(id),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	if book.Id > 0 {

		//调用dao updateBooks
		err = dao.UpdateBook(book)
		if err != nil {
			log.Println("update book is fail :", err)
		}
	} else {
		//调用dao addBooks
		err = dao.AddBook(book)
		if err != nil {
			log.Println("update book is fail :", err)
		}
	}
	//调用getbooks处理器函数 再查询一次数据库，
	GetPageBooks(w, r)

}

// 分页结论：  假设当前页面是pageNo， 每页显示条数
