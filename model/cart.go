package model

//cart 购物车

type Cart struct {
	CartId      string      //购物车id
	CartItems   []*CartItem //购物车中所有的购物项
	TotalCount  int64       // 购物车中图书的总数量，通过计算得到
	TotalAmount float64     //购物车中图书的总金额，通过计算得到
	UserId      int         // 购物车用户id
	UserName    string
}

// 获取购物车中图书的总数量
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	//遍历购物车中的购物项切片
	for _, value := range cart.CartItems {
		totalCount = totalCount + value.Count
	}
	return totalCount
}

// 获取购物车中图书的总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	//遍历购物车中的购物项切片
	for _, cartItem := range cart.CartItems {
		totalAmount += cartItem.GetAmount()
	}
	return totalAmount
}
