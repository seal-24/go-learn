package main

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"
	"time"
)

//查询单行
func QueryOneWait(DB *sql.DB,group *sync.WaitGroup) {
	user := new(User)   //用new()函数初始化一个结构体对象
	row := DB.QueryRow("select id,name from User where id=?", 1)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.id,&user.name); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	group.Done()
	//fmt.Println("Single row data:", *user)
}


func TestQueryOne(t *testing.T) {
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


		wg := &sync.WaitGroup{}
		wg.Add(3)
		for i := 0; i < 3; i++ {
			go QueryOneWait(DB,wg)
		}
		wg.Wait()



}


func BenchmarkQueryOne(b *testing.B) {
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

	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		wg := &sync.WaitGroup{}
		wg.Add(3)
		for i := 0; i < 3; i++ {
			go QueryOneWait(DB,wg)
		}
		wg.Wait()

	}

}

