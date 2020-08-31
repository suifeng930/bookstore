package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"html/template"
	"log"
	"net/http"
	"time"
)

func CheckOut(w http.ResponseWriter, r *http.Request) {

	//1. 获取session
	_, session := dao.IsLogin(r)
	//2.获取到用户登录信息
	userId := session.UserId
	cart, err := dao.GetCartByUserId(userId)
	if err != nil {
		log.Println(" 获取购物车失败：", err.Error())
	}

	//构造订单项
	orderId := utils.CreateUUID()
	dataString := time.Now().Format("2006-01-02 15:04:05")
	order := &model.Order{
		OrderId:     orderId,
		CreateTime:  dataString,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserId:      int64(userId),
	}
	//新增订单
	err = dao.AddOrder(order)
	if err != nil {
		log.Println("添加订单信息失败，", err.Error())
	}
	//保存订单项
	cartItmes := cart.CartItems
	for _, value := range cartItmes {
		//创建订单项
		orderItem := &model.OrderItem{
			Count:   value.Count,
			Amount:  value.Amount,
			Title:   value.Book.Title,
			Author:  value.Book.Author,
			Price:   value.Book.Price,
			ImgPath: value.Book.ImgPath,
			OrderId: orderId,
		}
		//添加到购物项
		err = dao.AddOrderItems(orderItem)
		if err != nil {
			log.Println(" 添加购物项失败；", err.Error())
		}
		//更新图书信息
		book := value.Book
		book.Sales = book.Sales + int(value.Count)
		book.Stock = book.Stock - int(value.Count)
		err = dao.UpdateBook(book)
		if err != nil {
			log.Println("更新图书信息失败：", err.Error())
		}

	}
	//清空购物车
	err = dao.DeleteCartByCartId(cart.CartId)
	if err != nil {
		log.Println(" 清空购物车失败 ；", err.Error())
	}
	session.OrderId = orderId
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	err = t.Execute(w, session)
	if err != nil {
		log.Println("解析模板失败：", err.Error())
	}

}
