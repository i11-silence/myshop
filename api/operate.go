package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"myshop/dao"
	"myshop/model"
	"myshop/service"
	"strconv"
)

func Praise(h *server.Hertz) {
	h.PUT("/comment/praise", func(ctx context.Context, c *app.RequestContext) {
		username, err := service.ValidateToken(c)
		if err != nil {
			fmt.Printf("validate token fail,err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		if username == "" {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		model := c.PostForm("model")
		res1, err := strconv.Atoi(model)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		commentId := c.PostForm("comment_id")
		res2, err := strconv.Atoi(commentId)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		if res1 == 1 {
			err = dao.Praise(res2, username)
			if err != nil {
				c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
				return
			}

		}
		if res1 == 2 {
			err = dao.Criticize(res2, username)
			if err != nil {
				c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
				return
			}
		}
		c.JSON(200, utils.H{"status": 10000, "info": "success"})
	})
}

func Order(h *server.Hertz) {
	h.POST("/operate/order", func(ctx context.Context, c *app.RequestContext) {
		username, err := service.ValidateToken(c)
		if err != nil {
			fmt.Printf("validate token fail,err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		if username == "" {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		var order model.Order
		err = c.Bind(&order)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		jsonData, err := json.Marshal(order.Order)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		res := string(jsonData)
		order_id, err := dao.InsertOrder(res, order.Username, order.Address, order.Total)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		c.JSON(200, utils.H{"status": 30000, "info": "success", "data": order_id})
	})
}
