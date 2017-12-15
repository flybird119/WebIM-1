package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	Id        int64  `orm:"auto"`
	Name      string `orm:"size(100)"`
	Password  string
	Beizhu    string
	Sex       string
	CreatTime time.Time
}

type Record struct {
	Id        int64  `orm:"auto"`
	Name      string `orm:"size(100)"`
	Context   string
	CreatTime time.Time
}

type Res struct {
	Name     string
	Password string
}

func RegisterDB() {
	//注册 model
	orm.RegisterModel(new(Record))
	//注册 model
	orm.RegisterModel(new(User))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库 密码不为空格式
	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8")

}

//保存聊天记录到mysql
func Insert(event Event) {
	//使用这个对象进行增删改查
	o := orm.NewOrm()
	record := new(Record)
	record.Name = event.User
	record.Context = event.Content
	record.CreatTime = time.Now()

	o.Insert(record)

}

//在user表中加入用户信息
func AddUserMsg(name string) {
	o := orm.NewOrm()
	user := new(User)
	user.Name = name
	user.CreatTime = time.Now()
	_, err := o.Insert(user)
	fmt.Println(err)

}

//查询数据库中是否包含用户名
func Query(name string) error {
	//res := new(Res)
	o := orm.NewOrm()
	user := User{Name: name}
	return o.Read(&user, "Name")
	//err := o.Raw("SELECT * FROM user WHERE name = ?", name).QueryRow(&res)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
	//return *res
}

//查询数据库中该用户名发出的所有发言

func QueryRecord(name string) []Record{

	o := orm.NewOrm()
	var  rs []Record
	//r := new(Record)
	num, err := o.Raw("SELECT * FROM record WHERE name = ?", name).QueryRows(&rs)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("查询不报错")
		fmt.Println(rs)
		fmt.Println(num)

	}
return rs
}
