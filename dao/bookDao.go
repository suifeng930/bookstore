package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"log"
	"strconv"
)

//get booke  获取数据库中说有的图书

func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err

	}
	var books []*model.Book
	for rows.Next() {
		// 一定要使用 类型断言
		book := &model.Book{}
		//给book 赋值
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			log.Println("book 赋值失败")
		}
		books = append(books, book)

	}
	return books, nil

}

// addbook  向数据库总添加一本图书
func AddBook(book *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"

	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
	if err != nil {
		log.Println("book dao insert into book is fail ", err.Error())
		return err
	}
	return nil

}

// delete book  删除一本图书
func DeleteBook(bookId int) error {
	sqlStr := "delete from books where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, bookId)
	if err != nil {
		log.Println("delete book is fails ", err.Error())
		return err

	}
	return nil
}

//get book ById
func GetBookById(bookId int) (*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id= ?"
	row := utils.Db.QueryRow(sqlStr, bookId)
	// 创建一个book
	book := &model.Book{}
	scan := row.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	if scan != nil {
		return nil, scan
	}
	return book, nil

}

//更新图书信息
func UpdateBook(book *model.Book) error {

	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"

	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.Id)
	if err != nil {
		return err
	}
	return nil

}

func GetPageBooks(pageNo string) (*model.Page, error) {

	PageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取数据库中图书的总记录数
	sqlStr := "select count(1) from books"
	//
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr)
	err := row.Scan(&totalRecord)
	if err != nil {
		log.Println("获取数据库表的总条数失败： ", err)
	}

	//设置每页显示4条记录
	var pageSize int64 = 4
	//设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	//获取当前页的图书
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"

	rows, err := utils.Db.Query(sqlStr2, (PageNo-int64(1))*pageSize, pageSize)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	//创建page

	page := &model.Page{
		Books:       books,
		PageNo:      PageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil

}

func GetPageBooksByPrice(pageNo string, minPrice, maxPrice string) (*model.Page, error) {

	PageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取数据库中图书的总记录数
	sqlStr := "select count(1) from books where  price between  ? and ?"
	//
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	err := row.Scan(&totalRecord)
	if err != nil {
		log.Println("获取数据库表的总条数失败： ", err)
	}

	//设置每页显示4条记录
	var pageSize int64 = 4
	//设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	//获取当前页的图书
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books  where  price between  ? and ? limit ?,?"

	rows, err := utils.Db.Query(sqlStr2, minPrice, maxPrice, (PageNo-int64(1))*pageSize, pageSize)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	//创建page

	page := &model.Page{
		Books:       books,
		PageNo:      PageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil

}
