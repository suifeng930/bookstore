package model

//订单表
type Order struct {
	OrderId     string  //订单号
	CreateTime  string  //生成订单时间
	TotalCount  int64   //订单总数量
	TotalAmount float64 //订单总金额
	State       int64   //订单状态   0 未发货   1 已发货  2 交易完成
	UserId      int64   //订单所属用户
}

//NoSend 未发货
func (order *Order) NoSend() bool {
	return order.State == 0
}

//SendComplate 已发货
func (order *Order) SendComplate() bool {
	return order.State == 1
}

//Complate 交易完成
func (order *Order) Complate() bool {
	return order.State == 2
}
