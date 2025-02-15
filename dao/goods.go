package dao

import (
	"fmt"
	"myshop/model"
)

func QueryGoods(username string) (goods []model.Good, err error) {
	sqlStr := `select * from goods `
	rows, err := Db.Query(sqlStr)
	if err != nil {
		fmt.Println("Query Good err:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var good model.Good
		err = rows.Scan(&good.Id, &good.Name, &good.Description, &good.Type, &good.Price, &good.Cover, &good.Link, &good.Publish_time)
		if err != nil {
			fmt.Println("Query Good err:", err)
			return nil, err
		}
		goods = append(goods, good)
	}
	for i, _ := range goods {
		sqlStr = `select count(*) from comment where good_id=?`
		err = Db.QueryRow(sqlStr, goods[i].Id).Scan(&goods[i].Comment_num)
		if err != nil {
			fmt.Println("Query Good err:", err)
			return nil, err
		}
	}
	for i, _ := range goods {
		sqlStr = `select count(*) from shopcart where good_id=? and username=?`
		err = Db.QueryRow(sqlStr, goods[i].Id, username).Scan(&goods[i].Is_addedCart)
		if err != nil {
			fmt.Println("Query Good err:", err)
			return nil, err
		}
	}
	return goods, nil
}

func SearchGoods(username string, productName string) (good model.Good, err error) {
	sqlStr := `select * from goods where name=?`
	err = Db.QueryRow(sqlStr, productName).Scan(&good.Id, &good.Name, &good.Description, &good.Type, &good.Price, &good.Cover, &good.Link, &good.Publish_time)
	if err != nil {
		fmt.Println("Query Good err:", err)
		return good, err
	}
	sqlStr = `select count(*) from comment where good_id=?`
	err = Db.QueryRow(sqlStr, good.Id).Scan(&good.Comment_num)
	if err != nil {
		fmt.Println("Query Good err:", err)
		return good, err
	}
	sqlStr = `select count(*) from shopcart where good_id=? and username=?`
	err = Db.QueryRow(sqlStr, good.Id, username).Scan(&good.Is_addedCart)
	if err != nil {
		fmt.Println("Query Good err:", err)
		return good, err
	}
	return good, nil
}

func InsertCart(username string, good_id int) error {
	//	sqlStr := `SELECT COUNT(*) from shopcart where good_id=?`
	sqlStr := `insert into shopcart(username,good_id) values(?,?)`
	_, err := Db.Exec(sqlStr, username, good_id)
	return err
}

func QueryCartList(username string) (goods []model.Good, err error) {
	sqlStr := `select g.* from goods g where g.id in(select good_id from shopcart where username=?)`
	rows, err := Db.Query(sqlStr, username)
	if err != nil {
		fmt.Println("Query Good err:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var good model.Good
		err = rows.Scan(&good.Id, &good.Name, &good.Description, &good.Type, &good.Price, &good.Cover, &good.Link, &good.Publish_time)
		if err != nil {
			fmt.Println("Query Good err:", err)
			return
		}
		goods = append(goods, good)
	}
	for i, _ := range goods {
		sqlStr = `select count(*) from comment where good_id=?`
		err = Db.QueryRow(sqlStr, goods[i].Id).Scan(&goods[i].Comment_num)
		if err != nil {
			fmt.Println("Query Good err:", err)
			return
		}
	}
	return goods, nil
}

func QueryGood(username string, id int) (good model.Good, err error) {
	sqlStr := `select * from goods where id=?`
	err = Db.QueryRow(sqlStr, id).Scan(&good.Id, &good.Name, &good.Description, &good.Type, &good.Price, &good.Cover, &good.Link, &good.Publish_time)
	if err != nil {
		fmt.Println("Query Good err:", err)
		return good, err
	}
	sqlStr = `select count(*) from comment where good_id=?`
	err = Db.QueryRow(sqlStr, id).Scan(&good.Comment_num)
	if err != nil {
		fmt.Println("Query Good err:", err)
		return good, err
	}
	sqlStr = `select count(*) from shopcart where good_id=? and username=?`
	err = Db.QueryRow(sqlStr, id, username).Scan(&good.Is_addedCart)
	if err != nil {
		fmt.Println("Query Good err:", err)
		return good, err
	}
	return good, nil
}

func QueryTypeGood(username string, goodType string) (goods []model.Good, err error) {
	sqlStr := `select * from goods where type=?`
	rows, err := Db.Query(sqlStr, goodType)
	if err != nil {
		fmt.Println("Query Good err:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var good model.Good
		err = rows.Scan(&good.Id, &good.Name, &good.Description, &good.Type, &good.Price, &good.Cover, &good.Link, &good.Publish_time)
		if err != nil {
			fmt.Println("Query Good err:", err)
			return nil, err
		}
		goods = append(goods, good)
	}
	for i, _ := range goods {
		sqlStr = `select count(*) from comment where good_id=?`
		err = Db.QueryRow(sqlStr, goods[i].Id).Scan(&goods[i].Comment_num)
		if err != nil {
			fmt.Println("Query Good err:", err)
			return nil, err
		}
	}
	for i, _ := range goods {
		sqlStr = `select count(*) from shopcart where good_id=? and username=?`
		err = Db.QueryRow(sqlStr, goods[i].Id, username).Scan(&goods[i].Is_addedCart)
		if err != nil {
			fmt.Println("Query Good err:", err)
			return nil, err
		}
	}
	return goods, nil
}
