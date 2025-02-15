package model

import "time"

type Comment struct {
	Id                  int    `json:"id"`
	Good_id             int    `json:"good_id"`
	Content             string `json:"content"`
	Publish_time        time.Time
	Publish_time_string string `json:"publish_time"`
	Username            string `json:"username"`
	Nickname            string `json:"nickname"`
	PraiseCount         int    `json:"praise_count"`
	IsPraised           int    `json:"is_praised"`
}
