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
	session.Order = order
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	err = t.Execute(w, session)
	if err != nil {
		log.Println("解析模板失败：", err.Error())
	}

}

//getorders  获取数据库中所有的订单
func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := dao.GetOrders()
	if err != nil {
		log.Println(" 获取数据库失败", err.Error())
	}
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))

	err = t.Execute(w, orders)
	if err != nil {
		log.Println(" 解析模板失败：", err.Error())

	}
}

//获取订单详细
func GetOrderItemByOrderId(w http.ResponseWriter, r *http.Request) {

	//获取订单id
	orderId := r.FormValue("orderId")
	//获取订单项
	orderItems, _ := dao.GetOrderItemsByOrderrId(orderId)
	log.Println(orderItems)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	err := t.Execute(w, orderItems)
	if err != nil {
		log.Println(" 解析模板失败：", err.Error())

	}
}

//根据用户id 获取订单详情
func GetMyOrders(w http.ResponseWriter, r *http.Request) {

	_, session := dao.IsLogin(r)
	//获取用户id
	userId := session.UserId
	//获取订单列表
	orders, err := dao.GetOrdersByUserId(userId)
	if err != nil {
		log.Println("获取订单失败", err.Error())
	}
	session.Orders = orders

	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	err = t.Execute(w, session)
	if err != nil {
		log.Println(" 解析模板失败：", err.Error())

	}

}

// SendOrder 发货
func SendOrder(w http.ResponseWriter, r *http.Request) {

	orderId := r.FormValue("orderId")

	err := dao.UpdateOrderStateByOrderId(orderId, 1)
	if err != nil {
		log.Println("更新订单状态信息失败：", err)
	}
	GetOrders(w, r)
}

// SendOrder 发货
func TakeOrder(w http.ResponseWriter, r *http.Request) {

	orderId := r.FormValue("orderId")

	err := dao.UpdateOrderStateByOrderId(orderId, 2)
	if err != nil {
		log.Println("更新订单状态信息失败：", err)
	}
	GetMyOrders(w, r)
}
