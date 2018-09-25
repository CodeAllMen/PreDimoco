package models

import (
	"github.com/astaxie/beego/orm"
)

type App struct {
	Id      int64  `orm:"pk;auto"`
	AppName string `orm:"size(50)"`
	Install int64
}

func init() {
	orm.RegisterModel(new(App))
}
