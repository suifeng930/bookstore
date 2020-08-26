package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"log"
)

//添加订单项
func AddOrderItems(item *model.OrderItem) error {
	sql := "insert into  order_items (count,amount,title,author,price,img_path,order_id) values (?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, item.Count, item.Amount, item.Title, item.Author, item.Price, item.ImgPath, item.OrderId)
	if err != nil {
		log.Println("insert into orderItems is fail ", err.Error())
		return err
	}
	return nil

}
