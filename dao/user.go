package dao

import (
	"database/sql"
	"fmt"
	"myshop/model"
	"time"
)

var Db *sql.DB

func InitDb() (err error) {
	dsn := "root:dly050425@tcp(127.0.0.1:3306)/myshop?charset=utf8&parseTime=True"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open db fail,err:", err)
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println("ping db fail,err:", err)
		return err
	}
	return err
}

func Insert(username string, password string) (err error) {
	sqlStr := `insert into users(username,password) values(?,?)`
	_, err = Db.Exec(sqlStr, username, password)
	return err
}

func ValidateCredentials(username string, password string) (bool, error) {
	sql := `select count(*) from users where username=? and password=?`
	var count int
	err := Db.QueryRow(sql, username, password).Scan(&count)
	return count > 0, err
}

func UpdatePassword(username string, newPassword string) (err error) {
	sql := `update users set password=? where username=?`
	_, err = Db.Exec(sql, newPassword, username)
	return err
}

func UpdateNickname(new string, username string) (err error) {
	sql := `update users set nickname=? where username=?`
	_, err = Db.Exec(sql, new, username)
	return err
}

func UpdateGender(new string, username string) (err error) {
	sql := `update users set gender=? where username=?`
	_, err = Db.Exec(sql, new, username)
	return err
}

func UpdatePhone(new string, username string) (err error) {
	sql := `update users set phone=? where username=?`
	_, err = Db.Exec(sql, new, username)
	return err
}

func UpdateBirthday(new time.Time, username string) (err error) {
	sql := `update users set birthday=? where username=?`
	_, err = Db.Exec(sql, new, username)
	return err
}

func QueryUser(username string) (user model.ReUser, err error) {
	sql := `select nickname,gender,phone,birthday from users where username=?`
	err = Db.QueryRow(sql, username).Scan(&user.Nickname, &user.Gender, &user.Phone, &user.Birthday)
	return user, err
}
