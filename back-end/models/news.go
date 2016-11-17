package models

import "github.com/astaxie/beego/orm"

type News struct {
	Id         int `orm:"auto;serial"`
	Title   string `orm:"size(30);null"`
	Content   string `orm:"size(20);null"json:"-"`
	Userid   int `orm:"size(20);null"`
	Rate   int `orm:"size(2);null"`
	Status int `orm:"size(2);null"`
}

func init() {
	orm.RegisterModel(new(News))
}
