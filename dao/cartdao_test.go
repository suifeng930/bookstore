package dao

import (
	"bookstore/model"
	"fmt"
	"log"
	"testing"
)

func TestSaveCart(t *testing.T) {

	fmt.Println("测试购物车")
	//t.Run("测试添加购物车", TestAddCart)
	t.Run("通过bookId 获取cartItem", TestGetCartItemByBookId)
	t.Run("通过cartId 获取cart", TestGetCartItemByCartId)

}

func TestAddCart(t *testing.T) {
	//设置要买的书
	book2 := &model.Book{
		Id:    34,
		Price: 27.2,
	}
	book1 := &model.Book{
		Id:    35,
		Price: 23.2,
	}
	//创建两个购物项
	cartItem1 := &model.CartItem{
		CartItemId: 0,
		Book:       book1,
		Count:      10,
		CartId:     "666888",
	}
	cartItem2 := &model.CartItem{
		CartItemId: 0,
		Book:       book2,
		Count:      10,
		CartId:     "666888",
	}

	var cartItems []*model.CartItem
	cartItems = append(cartItems, cartItem1)
	cartItems = append(cartItems, cartItem2)
	//创建购物车
	cart := &model.Cart{
		CartId:    "666888",
		CartItems: cartItems,
		UserId:    1,
	}
	err := AddCart(cart)
	if err != nil {
		log.Println(" 添加购物车失败，", err.Error())

	}
	log.Println(" 创建购物车成功")

}

func TestGetCartItemByBookId(t *testing.T) {

	bookId := "34"

	item, err := GetCartItemByBookId(bookId)
	if err != nil {
		log.Println("获取购物项失败")
	}
	log.Println(item)

}

func TestGetCartItemByCartId(t *testing.T) {

	bookId := "666888"

	item, err := GetCartItemByCartId(bookId)
	if err != nil {
		log.Println("获取购物项失败")
	}
	for key, value := range item {

		log.Printf("第%v 个购物项是： %v \n", key+1, value)

	}

}
