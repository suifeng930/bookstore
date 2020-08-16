package model

//购物项
type CartItem struct {
	CartItemId int64   //购物项id
	Book       *Book   //购物项中的图书信息
	Count      int64   //购物项中图书的数量
	Amount     float64 // 购物项中图书的金额小计，通过计算得到
	CartId     string  //购物车Id
}

//GetAmou
func (cartItem *CartItem) GetAmount() float64 {
	book := cartItem.Book
	price := book.Price
	return float64(cartItem.Count) * price

}
