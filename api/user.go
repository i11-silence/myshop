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
	"time"
)

func Register(h *server.Hertz) {
	h.POST("/user/register", func(ctx context.Context, c *app.RequestContext) {
		var u model.User
		err := c.Bind(&u)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		err = dao.Insert(u.Username, u.Password)
		if err != nil {
			fmt.Printf("insert fail,err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		c.JSON(200, utils.H{"status": 10000, "info": "success"})
	})
}

func Login(h *server.Hertz) {
	h.POST("/user/token", func(ctx context.Context, c *app.RequestContext) {
		var u model.User
		err := c.Bind(&u)
		if err != nil {
			fmt.Println("bind fail,err:", err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		var isvalid bool
		isvalid, err = dao.ValidateCredentials(u.Username, u.Password)
		if err != nil {
			fmt.Printf("validate fail,err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		if isvalid {
			token, err := service.GenerateToken(u.Username)
			refreshToken, err := service.GenerateRefresh_Token(u.Username)
			if err != nil {
				fmt.Printf("generate token fail,err:%v\n", err)
				c.JSON(400, utils.H{"status": 30000, "info": "failed"})
				return
			}
			c.JSON(200, utils.H{"status": 10000, "info": "success", "data": utils.H{"token": token, "refresh_token": refreshToken}})
		} else {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
		}
	})
}

func Refresh(h *server.Hertz) {
	h.GET("/user/token/refresh", func(ctx context.Context, c *app.RequestContext) {
		username, err := service.ValidaterefreshToken(c)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		if username == "" {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		token, err := service.GenerateToken(username)
		refreshToken, err := service.GenerateRefresh_Token(username)
		if err != nil {
			fmt.Printf("generate token fail,err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		c.JSON(200, utils.H{"status": 10000, "info": "success", "data": utils.H{"token": token, "refresh_token": refreshToken}})
	})
}

func ProfilePassword(h *server.Hertz) {
	h.PUT("/user/password", func(ctx context.Context, c *app.RequestContext) {
		username, err := service.ValidateToken(c)
		if err != nil {
			fmt.Printf("validate token fail,err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		if username == "" {
			fmt.Println("username is empty")
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		var p model.Password
		err = c.Bind(&p)
		if err != nil {
			fmt.Printf("bind fail,err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		isvalid, err := dao.ValidateCredentials(username, p.OldPassword)
		if err != nil {
			fmt.Printf("validate fail,err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		if isvalid {
			err = dao.UpdatePassword(username, p.NewPassword)
			if err != nil {
				fmt.Printf("update fail,err:%v\n", err)
				c.JSON(400, utils.H{"status": 30000, "info": "failed"})
				return
			}
			c.JSON(200, utils.H{"status": 10000, "info": "success"})
		} else {
			fmt.Printf("err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
		}
	})
}

func Update_profile(h *server.Hertz) {
	h.PUT("/user/info", func(ctx context.Context, c *app.RequestContext) {
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
		var u model.ReUser
		err = c.Bind(&u)
		if err != nil {
			fmt.Printf("bind fail,err:%v\n", err)
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		if u.Nickname != "" {
			err = dao.UpdateNickname(u.Nickname, username)
			if err != nil {
				fmt.Printf("update fail,err:%v\n", err)
				c.JSON(400, utils.H{"status": 30000, "info": "failed"})
				return
			}
		}
		if u.Gender != "" {
			err = dao.UpdateGender(u.Gender, username)
			if err != nil {
				fmt.Printf("update fail,err:%v\n", err)
				c.JSON(400, utils.H{"status": 30000, "info": "failed"})
				return
			}
		}
		if u.Phone != "" {
			err = dao.UpdatePhone(u.Phone, username)
			if err != nil {
				fmt.Printf("update fail,err:%v\n", err)
				c.JSON(400, utils.H{"status": 30000, "info": "failed"})
				return
			}
		}
		if u.Birthday != "" {
			layout := "2006-01-02"
			parsedTime, err := time.Parse(layout, u.Birthday)
			if err != nil {
				fmt.Printf("parse fail,err:%v\n", err)
				c.JSON(400, utils.H{"status": 30000, "info": "failed"})
				return
			}
			err = dao.UpdateBirthday(parsedTime, username)
			if err != nil {
				fmt.Printf("update fail,err:%v\n", err)
				c.JSON(400, utils.H{"status": 30000, "info": "failed"})
				return
			}
		}
		c.JSON(200, utils.H{"status": 10000, "info": "success"})
	})
}

func QueryUser(h *server.Hertz) {
	h.GET("/user/info/:user_id", func(ctx context.Context, c *app.RequestContext) {
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
		queryUsername := c.Param("user_id")
		if queryUsername == "" {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		var user model.ReUser
		user, err = dao.QueryUser(queryUsername)
		if err != nil {
			c.JSON(400, utils.H{"status": 30000, "info": "failed"})
			return
		}
		c.JSON(200, utils.H{"status": 10000, "info": "success", "data": user})
	})
}
