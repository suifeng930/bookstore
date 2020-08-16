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

// 根据用户id 从数据库中查询对应的购物项
//  1.先查询出cart记录，再查询封装cartItems
func GetCartByUserId(userId int) (*model.Cart, error) {

	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id=?"
	row := utils.Db.QueryRow(sqlStr, userId)
	cart := &model.Cart{}
	err := row.Scan(&cart.CartId, &cart.TotalCount, &cart.TotalAmount, &cart.UserId)
	if err != nil {

		log.Println("get cart by userId  is fails", err)
		return nil, err
	}
	//获取当前购物车里中 所有的购物项
	cartItems, err := GetCartItemByCartId(cart.CartId)
	if err != nil {
		log.Println("获取购物项失败")
		return nil, err
	}
	cart.CartItems = cartItems

	return cart, nil

}
