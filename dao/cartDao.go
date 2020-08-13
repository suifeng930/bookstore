package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"log"
)

// addCart  向购物车中插入购物车
func AddCart(cart *model.Cart) error {
	//
	sqlStr := "insert into carts (id,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cart.CartId, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserId)
	if err != nil {
		log.Println("add carts is fail :", err.Error())
		return err
	}
	//获取购物车中所有的购物项

	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		//将购物项插入到数据库中
		AddCartItem(cartItem)

	}
	return nil

}
