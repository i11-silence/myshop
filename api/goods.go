package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"myshop/dao"
	"myshop/model"
	"myshop/service"
	"strconv"
)

func QueryProductList(h *server.Hertz) {
	h.GET("/product/list", func(ctx context.Context, c *app.RequestContext) {
		username, _ := service.ValidateToken(c)
		goods, err := dao.QueryGoods(username)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		for i, good := range goods {
			goods[i].Publish_time_string = good.Publish_time.Format("Mon Jan 02 15:04:05 2006")
		}
		c.JSON(200, utils.H{"status": 10000, "info": "success", "data": goods})
	})
}

func SearchProductList(h *server.Hertz) {
	h.GET("/book/search", func(ctx context.Context, c *app.RequestContext) {
		username, _ := service.ValidateToken(c)
		productName := c.Query("product_name")
		good, err := dao.SearchGoods(username, productName)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		good.Publish_time_string = good.Publish_time.Format("Mon Jan 02 15:04:05 2006")
		c.JSON(200, utils.H{"status": 10000, "info": "success", "data": good})
	})
}

func AddCart(h *server.Hertz) {
	h.PUT("/product/addCart", func(ctx context.Context, c *app.RequestContext) {
		username, err := service.ValidateToken(c)
		if err != nil {
			fmt.Printf("validate token fail,err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		if username == "" {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		goodId := c.PostForm("product_id")
		id, err := strconv.Atoi(goodId)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		err = dao.InsertCart(username, id)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		c.JSON(200, utils.H{"status": 10000, "info": "success"})
	})
}

func QueryCartList(h *server.Hertz) {
	h.GET("/product/cart", func(ctx context.Context, c *app.RequestContext) {
		username, err := service.ValidateToken(c)
		if err != nil {
			fmt.Printf("validate token fail,err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		if username == "" {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		var goods []model.Good
		goods, err = dao.QueryCartList(username)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		for i, good := range goods {
			goods[i].Publish_time_string = good.Publish_time.Format("Mon Jan 02 15:04:05 2006")
		}
		for i, _ := range goods {
			goods[i].Is_addedCart = true
		}
		c.JSON(200, utils.H{"status": 10000, "data": goods})
	})
}

func QueryProduct(h *server.Hertz) {
	h.GET("/product/info/:product_id", func(ctx context.Context, c *app.RequestContext) {
		username, _ := service.ValidateToken(c)
		goodId := c.Param("product_id")
		id, err := strconv.Atoi(goodId)
		var good model.Good
		good, err = dao.QueryGood(username, id)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		good.Publish_time_string = good.Publish_time.Format("Mon Jan 02 15:04:05 2006")
		c.JSON(200, utils.H{"status": 10000, "info": "success", "data": good})
	})
}

func QueryTypeProduct(h *server.Hertz) {
	h.GET("/product/:type", func(ctx context.Context, c *app.RequestContext) {
		username, _ := service.ValidateToken(c)
		goodType := c.Param("type")
		goods, err := dao.QueryTypeGood(username, goodType)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		for i, good := range goods {
			goods[i].Publish_time_string = good.Publish_time.Format("Mon Jan 02 15:04:05 2006")
		}
		c.JSON(200, utils.H{"status": 10000, "info": "success", "data": goods})
	})
}
