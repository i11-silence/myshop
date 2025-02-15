package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Birthday string `json:"birthday"`
}

type Password struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
