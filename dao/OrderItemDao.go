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

//根据订单号来获取所有的订单项
func GetOrderItemsByOrderrId(orderId string) ([]*model.OrderItem, error) {
	//sql
	sqlStr := "select id,count,amount,title,author,price,img_path,order_id from order_items where  order_id=?"
	rows, err := utils.Db.Query(sqlStr, orderId)
	if err != nil {
		log.Println(" GetOrderItemsByOrderrId  get values is fail :", err.Error())
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		err := rows.Scan(&orderItem.OrderItemId, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderId)
		if err != nil {
			log.Println("sql rows scan is fail :", err.Error())
			return nil, err
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
