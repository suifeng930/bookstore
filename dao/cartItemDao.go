package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"log"
)

//   向购物项表中插入购物项
func AddCartItem(cartItem *model.CartItem) error {

	sqlStr := "insert into cart_items (count,amount,book_id,cart_id) values (?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		log.Println(" add cart_Items  is fail  :", err.Error())
		return err
	}
	return nil
}

// 根据bookId  获取对应的购物项
func GetCartItemByBookId(bookId string) (*model.CartItem, error) {

	sqlStr := "select id,count,amount,cart_id from cart_items where book_id=?"

	row := utils.Db.QueryRow(sqlStr, bookId)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartId)
	if err != nil {
		log.Println("get cartItem is fail ", err.Error())
		return nil, err
	}
	return cartItem, nil

}

//根据购物车id 查询所有的购物项
func GetCartItemByCartId(cartId string) ([]*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where cart_id=?"
	rows, err := utils.Db.Query(sqlStr, cartId)
	if err != nil {
		log.Println("get cart is fail ", err.Error())
		return nil, err
	}
	var cartItems []*model.CartItem
	cartItem := &model.CartItem{}
	for rows.Next() {

		err := rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartId)
		if err != nil {
			log.Println("get cartItem is fail ", err.Error())
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil

}
