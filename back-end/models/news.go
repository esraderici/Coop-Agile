package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
)

var missingAtribute = errors.New("Haber eklerken eksik bilgi var")


type News struct {
	Id      int `orm:"auto;serial"`
	Title   string `orm:"size(30);null"`
	Content string `orm:"size(20);null"json:"-"`
	Userid  int `orm:"size(20);null"`
	Rate    int `orm:"size(2);null"`
	Status  int `orm:"size(2);null"`
}

func init() {
	orm.RegisterModel(new(News))
}

//to save news
func (this *News) save() error {
	if (this.Content == "") || (this.Title == "" ){
		return missingAtribute
	}

	o := orm.NewOrm();
	o.Insert(this)

	return nil
}


//update
func (this *News) Update()  {
	if (this.Content == "") || (this.Title == "" ){
		return missingAtribute
	}


	o := orm.NewOrm();
	o.Update(this)

}
