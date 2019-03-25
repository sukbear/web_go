package main

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Userinfo struct {
	Uid        int `orm:"column(uid);pk"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username   string
	Departname string
	Created    time.Time
}

type User struct {
	Uid     int      `orm:"column(uid);pk"` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"` //设置一对多关系
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:mysql@/hrdb", 30)
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Userinfo), new(User), new(Profile), new(Tag), new(Post))
	// 创建 table
	orm.RunSyncdb("default", false, true)
}
func main() {
	// 插入数据
	o := orm.NewOrm()
	/*var user Userinfo
	user.Username = "zx11x"
	user.Departname = "zxx4x"
	user.Created = time.Now()
	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	} else {
		fmt.Println(err)
	}
	for  i:=1;i<10;i++{
		var u Userinfo
		u.Uid = i
		u.Username= "fdf"
		user.Departname = "zxx4x"
		u.Created= time.Now()
		id1, err := o.Insert(&u)
		if err == nil {
			fmt.Println(id1)
		} else {
			fmt.Println(err)
		}
	}*/

}
