package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {

	// 判断是否登录
	flag, session := dao.IsLogin(r)
	if flag { //已经登录

		//获取要添加的图书id
		bookId := r.FormValue("bookId")
		bookID, err := strconv.Atoi(bookId)
		if err != nil {
			log.Println("转换bookId失败", err.Error())
		}
		log.Println("要添加的图书id 是：", bookId)
		// 先根据用户的id查询购物车项；
		book, err := dao.GetBookById(bookID)
		log.Println("要添加的图书 是：", book)

		//获取用户id
		userId := session.UserId

		//判断数据库中是否存在当前用户的购物车
		cart, err := dao.GetCartByUserId(userId)
		log.Println("要添加的cart 是：", cart)
		if err != nil {
			log.Println(" 获取购物车失败")
		}
		if cart != nil {
			//当前用户已经存在购物车
			cartItem, err := dao.GetCartItemByBookIdAndCartId(bookId, cart.CartId)
			if err != nil {
				log.Println(" GetCartItemByBookIdAndCartId get cart book is fails ,", err.Error())
			}
			log.Println("要添加的cartItem 是：", cartItem)
			if cartItem != nil {
				//购物车的购物项已经存在，只需要更改购物项数量
				//1.获取购物车的所有购物项
				carts := cart.CartItems
				//2.遍历购物项
				log.Println("cartItem 是：", cartItem)
				for _, value := range carts {
					//3.找到当前的购物项
					log.Println("cartItem--》book 是：", cartItem.Book)
					if value.Book.Id == cartItem.Book.Id {
						//将当前购物项图书加1
						value.Count = value.Count + 1
						//4.更新数据库中该购物项的图书数量
						err := dao.UpdateBookCount(value)
						if err != nil {
							log.Println(" 更新购物车数量失败，", err)
						}
					}
				}
			} else {
				//购物车中的购物项还没有存在该图书，创建购物项
				Item := &model.CartItem{
					Book:   book,
					Count:  1,
					CartId: cart.CartId,
				}
				log.Println("添加购物车 数据：", Item)
				cart.CartItems = append(cart.CartItems, Item)
				err := dao.AddCartItem(Item)
				if err != nil {
					log.Println("创建购物项的数量，", err)
				}

			}
			// 不管购物车中是否存在图书对应的购物项，都需要更新购物车中的图书总数量和价格
			dao.UpdateCart(cart)

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
		w.Write([]byte("您刚刚将：" + book.Title + "添加到了购物车"))
	} else {
		w.Write([]byte("请先登录"))

	}

}

//根据用户的id 获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	//判断用户是否登录
	_, session := dao.IsLogin(r)
	//根据用户id
	userId := session.UserId
	//根据用户的idi 从数据看中获取对应的购物车
	cart, err := dao.GetCartByUserId(userId)
	if err != nil {
		log.Println(" getCart info 获取购物车失败", err.Error())
	}
	if cart != nil {
		//存在购物车
		//解析模板文件
		session.Cart = cart
		log.Println("cart is :", cart)
		log.Println("session is -->:", session)

		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		err := t.Execute(w, session)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		//用户还没有购物车
		log.Println("===============================")
		log.Println("session is -->:", session)
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		err := t.Execute(w, session)
		if err != nil {
			log.Println(err.Error())
		}
	}

}
