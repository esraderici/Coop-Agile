package models

import "github.com/astaxie/beego/orm"

type Comments struct {
	Id         int `orm:"auto;serial"`
	Userid   int `orm:"size(20);null"`
	Newsid   int `orm:"size(20);null"`
	Comment   string `orm:"size(40);null"`
}

func init() {
	orm.RegisterModel(new(Comments))
}
