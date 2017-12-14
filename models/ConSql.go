package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	_"github.com/go-sql-driver/mysql"
)

type Record struct {
	Id        int64  `orm:"auto"`
	Name      string `orm:"size(100)"`
	Context       string
	CreatTime time.Time
}
func RegisterDB() {
	//注册 model
	orm.RegisterModel(new(Record))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库 密码不为空格式
	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8")

}
