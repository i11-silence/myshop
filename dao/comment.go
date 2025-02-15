package dao

import (
	"encoding/json"
	"fmt"
	"myshop/model"
	"myshop/service"
	"time"
)

func QueryComments(good_id int, username string) ([]model.Comment, error) {
	var comments []model.Comment
	sqlStr := `select * from comment where good_id = ?`
	rows, err := Db.Query(sqlStr, good_id)
	if err != nil {
		fmt.Println("Query Comment err:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			comment                                   model.Comment
			praised_users_json, criticized_users_json string
			PraisedUsers, CriticizedUsers             []string
		)
		err = rows.Scan(&comment.Id, &comment.Good_id, &comment.Content, &comment.Publish_time, &comment.Username, &praised_users_json, &criticized_users_json)
		if err != nil {
			fmt.Println("Query Good err:", err)
			return nil, err
		}
		err = json.Unmarshal([]byte(praised_users_json), &PraisedUsers)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(criticized_users_json), &CriticizedUsers)
		if err != nil {
			return nil, err
		}
		comment.PraiseCount = len(PraisedUsers)
		if service.IsInSlice(username, PraisedUsers) {
			comment.IsPraised = 1
		} else if service.IsInSlice(username, CriticizedUsers) {
			comment.IsPraised = 2
		}

		sqlStr = `select nickname from users where username = ?`
		err = Db.QueryRow(sqlStr, comment.Username).Scan(&comment.Nickname)
		if err != nil {
			fmt.Println("Query Nickname err:", err)
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func InsertComment(username string, goodId int, content string) (int, error) {
	publishTime := time.Now()
	sqlStr := `insert into comment(good_id, content, publish_time,username,praised_users, criticized_users) values (?,?, ?,?,"[]","[]")`
	_, err := Db.Exec(sqlStr, goodId, content, publishTime, username)
	if err != nil {
		fmt.Println("Insert Comment err:", err)
		return 0, err
	}
	sqlStr = `select last_insert_id()`
	var commentId int
	err = Db.QueryRow(sqlStr).Scan(&commentId)
	if err != nil {
		fmt.Println("Insert Comment err:", err)
		return 0, err
	}
	return commentId, nil
}

func DeleteComment(commentId int, username string) error {
	sqlStr := `delete from comment where id = ? and username = ?`
	_, err := Db.Exec(sqlStr, commentId, username)
	return err
}

func UpdateComment(commentId int, content string, username string) error {
	sqlStr := `update comment set content = ? where id = ? and username = ?`
	_, err := Db.Exec(sqlStr, content, commentId, username)
	return err
}
