package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	_ "github.com/go-sql-driver/mysql"
	"myshop/api"
	"myshop/dao"
)

func main() {
	err := dao.InitDb()
	if err != nil {
		return
	}
	fmt.Println("连接数据库成功")

	defer dao.Db.Close()

	h := server.Default()

	api.Register(h)
	api.Login(h)
	api.Refresh(h)
	api.ProfilePassword(h)
	api.Update_profile(h)
	api.QueryUser(h)
	api.QueryProductList(h)
	api.SearchProductList(h)
	api.AddCart(h)
	api.QueryCartList(h)
	api.QueryProduct(h)
	api.QueryTypeProduct(h)
	api.QueryComment(h)
	api.AddComment(h)
	api.DeleteComment(h)
	api.UpdateComment(h)
	api.Praise(h)
	api.Order(h)

	h.Spin()
}
