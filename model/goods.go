package model

import "time"

type Good struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Type                string `json:"type"`
	Price               int    `json:"price"`
	Cover               string `json:"cover"`
	Link                string `json:"link"`
	Publish_time        time.Time
	Publish_time_string string `json:"publish_time"`
	Comment_num         int    `json:"comment_num"`
	Is_addedCart        bool   `json:"is_added_cart"`
}
