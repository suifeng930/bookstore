package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"log"
	"strconv"
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

// 根据bookId  cartId  获取对应的购物项
func GetCartItemByBookIdAndCartId(bookId string, cartId string) (*model.CartItem, error) {

	sqlStr := "select id,count,amount,cart_id from cart_items where book_id=? and cart_id=?"

	row := utils.Db.QueryRow(sqlStr, bookId, cartId)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartId)
	if err != nil {
		log.Println("get cartItem is fail ", err.Error())
		return nil, err
	}
	bookID, err := strconv.Atoi(bookId)
	book, err := GetBookById(bookID)
	cartItem.Book = book
	return cartItem, nil

}

// 根据图书的id 和购物车的id 以及图书的数量更新购物车项中图书的数量
func UpdateBookCount(item *model.CartItem) error {
	sqlStr := "update cart_items set count=?,amount=? where book_id=? and cart_id=?"
	_, err := utils.Db.Exec(sqlStr, item.Count, item.GetAmount(), item.Book.Id, item.CartId)
	if err != nil {
		log.Println("更新图书信息，by bookId and cartId ", err.Error())
		return err
	}
	return nil

}

//根据购物车id 查询所有的购物项
func GetCartItemByCartId(cartId string) ([]*model.CartItem, error) {
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id=?"
	rows, err := utils.Db.Query(sqlStr, cartId)
	if err != nil {
		log.Println("get cart is fail ", err.Error())
		return nil, err
	}
	var cartItems []*model.CartItem
	cartItem := &model.CartItem{}
	for rows.Next() {
		var bookId string

		err := rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartId)
		if err != nil {
			log.Println("get cartItem is fail ", err.Error())
			return nil, err
		}
		//根据bookId 获取图书信息
		bookID, err := strconv.Atoi(bookId)
		book, err := GetBookById(bookID)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil

}

// 根据购物车id 删除购物项
func DeleteCartItemByCartId(cartId string) error {
	sql := "delete from cart_items where cart_id=?"
	_, err := utils.Db.Exec(sql, cartId)
	if err != nil {
		log.Println(" delete by cartId is fail :", err.Error())
		return err
	}
	return nil
}
