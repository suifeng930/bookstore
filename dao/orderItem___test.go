package dao

import (
	"fmt"
	"log"
	"testing"
)

func TestAddOrderItems(t *testing.T) {

	fmt.Println("测试用户")
	t.Run("测试添加session", TestGetOrderItemsByOrderrId)
	//t.Run("删除session", TestDeleteSession)

}

func TestGetOrderItemsByOrderrId(t *testing.T) {
	orderId := "c750c2b8-1228-40d4-6bbe-2e68835c87dc"
	items, err := GetOrderItemsByOrderrId(orderId)
	if err != nil {
		log.Println("获取订单项失败")
	}
	for _, value := range items {

		fmt.Println("订单项：", value)
	}

}
