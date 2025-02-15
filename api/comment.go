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

func QueryComment(h *server.Hertz) {
	h.GET("/comment/:product_id", func(ctx context.Context, c *app.RequestContext) {
		username, _ := service.ValidateToken(c)
		goodID := c.Param("product_id")
		res, err := strconv.Atoi(goodID)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			fmt.Println(err.Error())
			return
		}
		comments, err := dao.QueryComments(res, username)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			fmt.Println(err.Error())
			return
		}
		for i, comment := range comments {
			comments[i].Publish_time_string = comment.Publish_time.Format("Mon Jan 02 15:04:05 2006")
		}
		c.JSON(200, utils.H{"status": 10000, "info": "success", "data": comments})
	})
}

func AddComment(h *server.Hertz) {
	h.POST("/comment/:product_id", func(ctx context.Context, c *app.RequestContext) {
		username, err := service.ValidateToken(c)
		if err != nil || username == "" {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		goodID := c.Param("product_id")
		res, err := strconv.Atoi(goodID)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		_, err = dao.QueryGood(username, res)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": "不能对不存在的商品进行评论"})
			return
		}
		var comment_content model.Comment
		err = c.Bind(&comment_content)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		commentId, err := dao.InsertComment(username, res, comment_content.Content)
		c.JSON(200, utils.H{"status": 10000, "info": "success", "data": commentId})
	})
}

func DeleteComment(h *server.Hertz) {
	h.DELETE("/comment/:comment_id", func(ctx context.Context, c *app.RequestContext) {
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
		commentId := c.Param("comment_id")
		res, err := strconv.Atoi(commentId)
		err = dao.DeleteComment(res, username)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		c.JSON(200, utils.H{"status": 10000, "info": "success"})
	})
}

func UpdateComment(h *server.Hertz) {
	h.PUT("/comment/:comment_id", func(ctx context.Context, c *app.RequestContext) {
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
		commentId := c.Param("comment_id")
		res, err := strconv.Atoi(commentId)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		var comment_content model.Comment
		err = c.Bind(&comment_content)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		err = dao.UpdateComment(res, comment_content.Content, username)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": err.Error()})
			return
		}
		c.JSON(200, utils.H{"status": 10000, "info": "success"})
	})
}
