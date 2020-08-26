package model

type OrderItem struct {
	OrderItemId int64   //订单项id
	Count       int64   //订单数量
	Amount      float64 //订单项中图书的金额小计
	Title       string  //订单项图书的书名
	Author      string  //作者
	Price       float64 //价格
	ImgPath     string  //订单图书封面
	OrderId     string  //订单项所属订单
}
