package main

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Info struct {
	Id int64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Path string `json:"path" db:"path"`
	Note string	`json:"note" db:"note"`
	CreateTime int64 `json:"create_time" db:"createTime"`
}

// 定义数据库
var Db *sqlx.DB

func init(){
	db,err := sqlx.Open("mysql","root:709463253@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(db,err)
	if err != nil{
		log.Fatal(err)
	}
	err = db.Ping() // 查看数据库是否连接
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("数据库已连接")
	Db = db
}

// 添加
func InfoAdd(mod *Info) error{
	result,err := Db.Exec("insert into info (name,path,note,createTime) values (?,?,?,?)",mod.Name,mod.Path,mod.Note,mod.CreateTime)
	fmt.Println(result,err,"---")
	if err!=nil{
		return err
	}
	id,_ := result.LastInsertId()
	if id <1 {
		return errors.New("添加失败")
	}
	return nil
}

// 删除
func InfoDelete(id int64) error{
	result,err := Db.Exec("delete from info where id = ?",id)
	if err !=nil{
		return err
	}
	rows,_:=result.RowsAffected()
	if rows!=1{
		return errors.New("删除失败")
	}
	return nil
}

//  InfoGet查询某个条件
func InfoGet(id int64) (Info,error){
	mod := Info{}
	err := Db.Get(&mod,"select * from info where id = ?",id)
	return mod,err
}

