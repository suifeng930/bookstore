package model

import "time"

//订单表
type Order struct {
	OrderId     string    //订单号
	CreateTime  time.Time //生成订单时间
	TotalCount  int64     //订单总数量
	TotalAmount float64   //订单总金额
	State       int64     //订单状态   0 未发货   1 已发货  2 交易完成
	UserId      int64     //订单所属用户
}
