package dao

import (
	"github.com/cloudwego/hertz/pkg/common/json"
	"myshop/service"
)

func Praise(commentId int, username string) error {
	var (
		praisedUsers, criticizedUsers             []string
		praised_users_json, criticized_users_json string
	)
	sqlStr := `select praised_users,criticized_users from comment where id = ?`
	err := Db.QueryRow(sqlStr, commentId).Scan(&praised_users_json, &criticized_users_json)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(praised_users_json), &praisedUsers)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(criticized_users_json), &criticizedUsers)
	if err != nil {
		return err
	}
	if service.IsInSlice(username, criticizedUsers) {
		newCriticizedUsers := service.DeleteInSlice(username, criticizedUsers)
		jsonData, err := json.Marshal(newCriticizedUsers)
		if err != nil {
			return err
		}
		criticized_users_json = string(jsonData)
		sqlStr = `update comment set criticized_users = ? where id = ?`
		_, err = Db.Exec(sqlStr, criticized_users_json, commentId)
		if err != nil {
			return err
		}
	}
	praisedUsers = append(praisedUsers, username)
	jsonData, err := json.Marshal(praisedUsers)
	if err != nil {
		return err
	}
	praised_users_json = string(jsonData)
	sqlStr = `update comment set praised_users = ? where id = ?`
	_, err = Db.Exec(sqlStr, praised_users_json, commentId)
	if err != nil {
		return err
	}
	return nil
}

func Criticize(commentId int, username string) error {
	var (
		praisedUsers, criticizedUsers             []string
		praised_users_json, criticized_users_json string
	)
	sqlStr := `select praised_users,criticized_users from comment where id = ?`
	err := Db.QueryRow(sqlStr, commentId).Scan(&praised_users_json, &criticized_users_json)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(praised_users_json), &praisedUsers)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(criticized_users_json), &criticizedUsers)
	if err != nil {
		return err
	}
	if service.IsInSlice(username, praisedUsers) {
		newPraisedUsers := service.DeleteInSlice(username, praisedUsers)
		jsonData, err := json.Marshal(newPraisedUsers)
		if err != nil {
			return err
		}
		praised_users_json = string(jsonData)
		sqlStr = `update comment set praised_users = ? where id = ?`
		_, err = Db.Exec(sqlStr, praised_users_json, commentId)
		if err != nil {
			return err
		}
	}
	criticizedUsers = append(criticizedUsers, username)
	jsonData, err := json.Marshal(criticizedUsers)
	if err != nil {
		return err
	}
	criticized_users_json = string(jsonData)
	sqlStr = `update comment set criticized_users = ? where id = ?`
	_, err = Db.Exec(sqlStr, criticized_users_json, commentId)
	if err != nil {
		return err
	}
	return nil
}

func InsertOrder(order string, username string, Address string, total float64) (int, error) {
	sqlStr := `insert into order(username,address,order,total) values(?,?,?,?)`
	_, err := Db.Exec(sqlStr, username, Address, order, total)
	if err != nil {
		return 0, err
	}
	sqlStr = `select last_insert_id()`
	var orderId int
	err = Db.QueryRow(sqlStr).Scan(&orderId)
	if err != nil {
		return 0, err
	}
	return orderId, nil
}
