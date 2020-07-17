package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

//连接数据库
func init() {
	var err error
	//Db, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名?charset=编码")
	Db, err = sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mall?charset=utf8")
	if err != nil {
		panic(err)
	}
}

type UmsRole struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	AdminCount  int    `db:"admin_count"`
	CreateTime  string `db:"create_time"`
	Status      int    `db:"status"`
	Sort        int    `db:"sort"`
}

func main() {

	//查询多条
	var list []UmsRole
	err := Db.Select(&list, "SELECT name,status,description FROM ums_role where name like ?", "%商品%")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)

}
