package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"log"
	"net/http"
	"strconv"
)

//添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {

	//获取要添加的图书id
	bookId := r.FormValue("bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		log.Println("转换bookId失败", err.Error())
	}
	log.Println("要添加的图书id 是：", bookId)
	// 先根据用户的id查询购物车项；
	book, err := dao.GetBookById(id)
	// 判断是否登录
	_, session := dao.IsLogin(r)
	//获取用户id
	userId := session.UserId

	//判断数据库中是否存在当前用户的购物车
	cart, err := dao.GetCartByUserId(userId)
	if err != nil {
		log.Println(" 获取购物车失败")
	}
	if cart != nil {
		//当前用户已经存在购物车
	} else {
		//没有存在购物车，添加购物车
		//1.创建一个购物项
		cartId := utils.CreateUUID()
		cart := &model.Cart{
			CartId: cartId,
			UserId: userId,
		}
		var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			Book:   book,
			Count:  1,
			CartId: cartId,
		}
		cartItems = append(cartItems, cartItem)
		cart.CartItems = cartItems
		log.Println("添加购物车 数据：", cart)
		err := dao.AddCart(cart)
		if err != nil {
			log.Println("创建购物车失败，", err)
		}
	}

}
