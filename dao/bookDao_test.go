package dao

import (
	"bookstore/model"
	"fmt"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("测试bookdao 中的方法")
	m.Run()

}

func TestGetBooks(t *testing.T) {

	log.Println("测试book dao get books 方法")
	//t.Run("测试获取所有图书:", testGetBooks)
	//t.Run("测试添加图书:", testAddBooks)
	//t.Run("测试删除图书:", TestDeleteBooks)
	//t.Run("测试查询一本图书:", TestGetBookById)
	//t.Run("测试修改一本图书:", TestUpdateBook)
	t.Run("查询分页信息:", TestGetPageBooks)
}

func testGetBooks(t *testing.T) {
	books, err := GetBooks()
	if err != nil {
		log.Println("获取图书失败")

	}
	for key, value := range books {
		log.Printf("第 %v 本图书的信息是： %v\n", key+1, value)
	}
}

func testAddBooks(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.ipg",
	}
	addBook := AddBook(book)
	if addBook != nil {
		log.Println("add book is fail ")
	}

}

func TestDeleteBooks(t *testing.T) {
	bookId := 49
	book := DeleteBook(bookId)
	if book != nil {
		log.Println("delete book is fail ")
	}

}

func TestGetBookById(t *testing.T) {
	bookId := 17
	book, err := GetBookById(bookId)
	if err != nil {
		log.Println("delete book is fail ")
	}

	log.Println(book)

}

func TestUpdateBook(t *testing.T) {
	book := &model.Book{
		Id:     52,
		Title:  "小马哥",
		Author: "小马哥",
		Price:  45.69,
		Sales:  123,
		Stock:  12,
	}
	err := UpdateBook(book)
	if err != nil {
		log.Println("update book is fail ")
	}

	log.Println(book)

}

func TestGetPageBooks(t *testing.T) {

	page, Err := GetPageBooks("8")
	if Err != nil {
		log.Println(Err)
	}

	fmt.Println("当前页数是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.TotalRecord)
	fmt.Println("当前页 有图书是：")
	for _, value := range page.Books {
		fmt.Println("图书信息是：", value)

	}
}
