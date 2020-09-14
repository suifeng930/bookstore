package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"log"
)

// 添加订单
func AddOrder(order *model.Order) error {

	sql := "insert into orders (id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, order.OrderId, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserId)
	if err != nil {
		log.Println(" insert into order is fail :", err.Error())
		return err

	}
	return nil

}

// getOrders   获取所有的购物车
func GetOrders() ([]*model.Order, error) {
	spl := "select id,create_time,total_count,total_amount,state,user_id from orders  "
	rows, err := utils.Db.Query(spl)
	if err != nil {
		log.Println("getOrders is fail :", err.Error())
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		err := rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		if err != nil {
			log.Println("getOrders is fail ", err.Error())
		}
		orders = append(orders, order)
	}
	return orders, nil

}

//通过用户id 获取订单列表
func GetOrdersByUserId(userId int) ([]*model.Order, error) {
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders  where user_id=? "
	rows, err := utils.Db.Query(sqlStr, userId)
	if err != nil {
		log.Println("get orders by userId is fail ：", err.Error())
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		err := rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		if err != nil {
			log.Println("get orders is fail :", err.Error())
			return nil, err
		}
		orders = append(orders, order)
	}
	log.Println("orders:", orders)
	return orders, nil

}
