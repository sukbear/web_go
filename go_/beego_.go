package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type user struct {
	Id int
	Name string `orm:"size(100)"`
}

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:mysql@/hrdb", 30)
	//注册定义的model
	orm.RegisterModel(new(user))

	// 创建table
	orm.RunSyncdb("default", false, true)
}
func main() {
	o := orm.NewOrm()
	orm.SetMaxIdleConns("default", 30)
	user := user{1,"fuck"}
	id, err:= o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// 更新表
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := user{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// 删除表
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}