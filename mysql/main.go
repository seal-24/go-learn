package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	_ "net/http/pprof"
	"time"
)

//数据库连接信息
const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK = "tcp"
	SERVER = "127.0.0.1"
	PORT = 3306
	DATABASE = "test"
)

//user表结构体定义
type User struct {
	id int `json:"id" form:"id"`
	name string `json:"name" form:"name"`

}
//
//docker run -itd --name mysql-test -p 3307:3307 -e MYSQL_ROOT_PASSWORD=123456 mysql
//docker exec mysql-test  /bin/bash

func main() {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", conn)
	if err != nil {
		//fmt.Println("connection to mysql failed:", err)
		fmt.Println("connection to mysql failed:", err)
		return
	}

	DB.SetConnMaxLifetime(3 * time.Second)    //最大连接周期，超时的连接就close? not working
	DB.SetMaxOpenConns(1)                  //设置最大连接数
	//CreateTable(DB)
	//InsertData(DB)
	//go func() {
	//	for i := 1 ; i< 1000000; i++{
	//		QueryOne(DB)
	//		time.Sleep(time.Second)
	//
	//	}
	//}()

	// curl http://127.0.0.1:6060/debug/pprof/trace\?seconds\=20 > trace.out
	go http.ListenAndServe("0.0.0.0:6060", nil)
	for {
		for i:= 1; i < 4;i++ {
			go QueryOne(DB)
		}
		time.Sleep(time.Second)

	}


}


	//QueryMulti(DB)
	//UpdateData(DB)
	//DeleteData(DB)


func CreateTable(DB *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS users(
	id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
	username VARCHAR(64),
	password VARCHAR(64),
	status INT(4),
	createtime INT(10)
	); `

	if _, err := DB.Exec(sql); err != nil {
		fmt.Println("create table failed:", err)
		return
	}
	fmt.Println("create table successd")
}

//插入数据
//func InsertData(DB *sql.DB) {
//	result,err := DB.Exec("insert INTO users(username,password) values(?,?)","test","123456")
//	if err != nil{
//		fmt.Printf("Insert data failed,err:%v", err)
//		return
//	}
//	lastInsertID,err := result.LastInsertId()    //获取插入数据的自增ID
//	if err != nil {
//		fmt.Printf("Get insert id failed,err:%v", err)
//		return
//	}
//	fmt.Println("Insert data id:", lastInsertID)
//
//	rowsaffected,err := result.RowsAffected()  //通过RowsAffected获取受影响的行数
//	if err != nil {
//		fmt.Printf("Get RowsAffected failed,err:%v",err)
//		return
//	}
//	fmt.Println("Affected rows:", rowsaffected)
//}

//查询单行
func QueryOne(DB *sql.DB) {
	user := new(User)   //用new()函数初始化一个结构体对象
	row := DB.QueryRow("select id,name from User where id=?", 1)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.id,&user.name); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	//fmt.Println("Single row data:", *user)
}

////查询多行
//func QueryMulti(DB *sql.DB) {
//	user := new(User)
//	rows, err := DB.Query("select id,username,password from users where id = ?", 2)
//
//	defer func() {
//		if rows != nil {
//			rows.Close()   //关闭掉未scan的sql连接
//		}
//	}()
//	if err != nil {
//		fmt.Printf("Query failed,err:%v\n", err)
//		return
//	}
//	for rows.Next() {
//		err = rows.Scan(&user.id, &user.id, &user.Password)  //不scan会导致连接不释放
//		if err != nil {
//			fmt.Printf("Scan failed,err:%v\n", err)
//			return
//		}
//		fmt.Println("scan successd:", *user)
//	}
//}
//
////更新数据
//func UpdateData(DB *sql.DB){
//	result,err := DB.Exec("UPDATE users set password=? where id=?","111111",1)
//	if err != nil{
//		fmt.Printf("Insert failed,err:%v\n", err)
//		return
//	}
//	fmt.Println("update data successd:", result)
//
//	rowsaffected,err := result.RowsAffected()
//	if err != nil {
//		fmt.Printf("Get RowsAffected failed,err:%v\n",err)
//		return
//	}
//	fmt.Println("Affected rows:", rowsaffected)
//}
//
////删除数据
//func DeleteData(DB *sql.DB){
//	result,err := DB.Exec("delete from users where id=?",1)
//	if err != nil{
//		fmt.Printf("Insert failed,err:%v\n",err)
//		return
//	}
//	fmt.Println("delete data successd:", result)
//
//	rowsaffected,err := result.RowsAffected()
//	if err != nil {
//		fmt.Printf("Get RowsAffected failed,err:%v\n",err)
//		return
//	}
//	fmt.Println("Affected rows:", rowsaffected)
//}
